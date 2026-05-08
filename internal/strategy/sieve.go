package strategy

// Sieve implements PrimeStrategy using the Sieve of Eratosthenes algorithm.
type Sieve struct{}

// Name returns the identifier of this strategy.
func (s *Sieve) Name() string {
	return "sieve"
}

// Generate returns all prime numbers in the range [start, end] using sieve.
func (s *Sieve) Generate(start, end int) ([]int, error) {
	primes := make([]int, 0)

	// Edge case: no primes below 2
	if end < 2 {
		return primes, nil
	}

	// Step 1: Create a boolean array
	// Assume all numbers are prime initially
	isPrime := make([]bool, end+1)
	for i := 2; i <= end; i++ {
		isPrime[i] = true
	}

	// Step 2: Mark non-primes
	// Only go till sqrt(end)
	for i := 2; i*i <= end; i++ {
		if isPrime[i] {
			// Mark all multiples of i as false
			for j := i * i; j <= end; j += i {
				isPrime[j] = false
			}
		}
	}

	// Step 3: Collect primes in given range
	if start < 2 {
		start = 2
	}

	for i := start; i <= end; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}

	return primes, nil
}