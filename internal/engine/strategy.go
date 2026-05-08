package engine

// PrimeStrategy defines what every prime algorithm should implement
type PrimeStrategy interface {

	// returns name of the algorithm (used to identify it)
	Name() string

	// generates prime numbers between start and end
	Generate(start, end int) ([]int, error)
}
