package store

import (
	"sync"
	"time"
)

// Log represents one execution record of prime generation
type Log struct {
	Timestamp time.Time // when request was made
	Start     int       // start of range
	End       int       // end of range
	Algorithm string    // algorithm used
	TimeUs    int64     // execution time
	Count     int       // number of primes returned
}

// Store is an in-memory database.
// It uses a mutex to handle concurrent access safely.
type Store struct {
	mu   sync.Mutex
	logs []Log
}

// NewStore initializes a new Store
func NewStore() *Store {
	return &Store{
		logs: make([]Log, 0),
	}
}

// Add saves a new log entry
func (s *Store) Add(log Log) {
	s.mu.Lock()         // lock before writing
	defer s.mu.Unlock() 

	s.logs = append(s.logs, log)
}

// GetAll returns all logs 
func (s *Store) GetAll() []Log {
	s.mu.Lock()
	defer s.mu.Unlock()

	// return copy to avoid external modification
	result := make([]Log, len(s.logs))
	copy(result, s.logs)

	return result
}