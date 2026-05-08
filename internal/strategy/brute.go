package strategy

import "math"

// Brute uses simple approach to find primes
type Brute struct{}

func (b *Brute) Name() string {
	return "brute"
}

// Generate finds primes between start and end
func (b *Brute) Generate(start, end int) ([]int, error) {
	primes := make([]int, 0)

	// no primes below 2
	if end < 2 {          
		return primes, nil
	}

	// adjust start if needed
	if start < 2 {
		start = 2
	}

	// check each number
	for i := start; i <= end; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}

	return primes, nil
}

// checks if a number is prime
func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	// only check till sqrt(n)
	limit := int(math.Sqrt(float64(n)))

	for i := 2; i <= limit; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}