package service

import (
	"fmt"
	"prime-generator/internal/engine"
	"prime-generator/internal/store"
	"prime-generator/internal/strategy"
	"time"
)

//prime service act as the business logic layer.
//its sits beteen handler(API) and engine (core logic)\

type PrimeService struct{
	engine *engine.Engine
	store  *store.Store
}

func NewPrimeService() *PrimeService{
	e:=engine.NewEngine()
	st := store.NewStore()
	e.Register(&strategy.Brute{})
	e.Register(&strategy.Sieve{})


	return &PrimeService{
		engine: e,
		store:  st,
	}
}



func (s *PrimeService) Generate(start, end int, algo string) (map[string]interface{}, error) {

	if start > end {
		return nil, fmt.Errorf("start cannot be greater than end")
	}

	if start < 2 {
		start = 2
	}
	startTime := time.Now()


	if algo == "auto" {
		algo = s.selectAlgo(start, end)
	}

	size := end - start

// prevent memory explosion
if algo == "sieve" && size > 10_000_000 {
	return nil, fmt.Errorf("range too large for sieve")
}
	primes, err := s.engine.Execute(algo, start, end)
	if err != nil {
		return nil, err
	}

	duration := time.Since(startTime)

	
	s.store.Add(store.Log{
		Timestamp: time.Now(),
		Start:     start,
		End:       end,
		Algorithm: algo,
		TimeUs:    duration.Microseconds(), // changed
		Count:     len(primes),
	})

	return map[string]interface{}{
		"primes":    primes,
		"count":     len(primes),
		"algorithm": algo,
		"time_us":   duration.Microseconds(),
	}, nil
}
// selectAlgo decides best algorithm based on input size
func (s *PrimeService) selectAlgo(start, end int) string {
	size := end - start

	if size < 10000 {
		return "brute"
	} else if size < 10_000_000 {
		return "sieve"
	}

	return "brute"
}

func (s *PrimeService) GetStats() map[string]interface{} {

	logs := s.store.GetAll()

	totalRequests := len(logs)

	if totalRequests == 0 {
		return map[string]interface{}{
			"total_requests": 0,
		}
	}

	var totalTime int64

	type AlgoStats struct {
		count int
		total int64
		max   int64
		min   int64
	}

	statsMap := make(map[string]*AlgoStats)

	for _, log := range logs {
		totalTime += log.TimeUs

		if _, exists := statsMap[log.Algorithm]; !exists {
			statsMap[log.Algorithm] = &AlgoStats{
				max: log.TimeUs,
				min: log.TimeUs,
			}
		}

		stat := statsMap[log.Algorithm]

		stat.count++
		stat.total += log.TimeUs

		if log.TimeUs > stat.max {
			stat.max = log.TimeUs
		}

		if log.TimeUs < stat.min {
			stat.min = log.TimeUs
		}
	}

	algoResponse := make(map[string]interface{})

	for algo, stat := range statsMap {
		algoResponse[algo] = map[string]interface{}{
			"count":       stat.count,
			"avg_time_us": stat.total / int64(stat.count),
			"max_time_us": stat.max,
			"min_time_us": stat.min,
		}
	}

	return map[string]interface{}{
		"total_requests":       totalRequests,
		"overall_avg_time_us":  totalTime / int64(totalRequests),
		"algorithms":           algoResponse,
	}
}