package primes

import (
	"fmt"
	"testing"
	"time"
)

func TestPrime(t *testing.T) {

	primeNumbers := []primeCandidate{1000003, 1000033, 1000037, 1003279, 1999993, 982451653}

	start := time.Now()

	for _, number := range primeNumbers {
		if !number.isPrime() {
			t.Errorf("%d is prime\n", number)
		}
	}

	candidate := primeCandidate(534395689)
	if candidate.isPrime() {
		t.Errorf("%d is not a prime\n", candidate)
	}

	fmt.Printf("Basic test took %s\n", time.Since(start))

	start = time.Now()

	for _, number := range primeNumbers {
		if !number.isPrimeAKS() {
			t.Errorf("%d is prime\n", number)
		}
	}

	candidate = primeCandidate(534395689)
	if candidate.isPrimeAKS() {
		t.Errorf("%d is not a prime\n", candidate)
	}

	fmt.Printf("AKS test took %s\n", time.Since(start))

}
