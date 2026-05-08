package service

import "testing"

// basic test to check if prime generation works correctly
func TestGenerate_Basic(t *testing.T) {
	svc := NewPrimeService()

	result, err := svc.Generate(1, 10, "brute")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	primes := result["primes"].([]int)

	expected := []int{2, 3, 5, 7}

	if len(primes) != len(expected) {
		t.Fatalf("expected %v, got %v", expected, primes)
	}

	for i := range primes {
		if primes[i] != expected[i] {
			t.Fatalf("expected %v, got %v", expected, primes)
		}
	}
}

// checks if invalid range returns error
func TestGenerate_InvalidRange(t *testing.T) {
	svc := NewPrimeService()

	_, err := svc.Generate(10, 1, "brute")

	if err == nil {
		t.Fatalf("expected error but got nil")
	}
}

// checks if auto strategy selects correct algorithm
func TestGenerate_Auto(t *testing.T) {
	svc := NewPrimeService()

	result, _ := svc.Generate(1, 50000, "auto")

	algo := result["algorithm"].(string)

	if algo != "sieve" {
		t.Fatalf("expected sieve, got %s", algo)
	}
}

// checks if stats are recorded properly
func TestStats(t *testing.T) {
	svc := NewPrimeService()

	svc.Generate(1, 10, "brute")
	svc.Generate(1, 100, "sieve")

	stats := svc.GetStats()

	total := stats["total_requests"].(int)

	if total != 2 {
		t.Fatalf("expected 2, got %d", total)
	}
}