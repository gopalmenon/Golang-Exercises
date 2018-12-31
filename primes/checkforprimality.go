package primes

import "math"

type primeCandidate int64

func (p primeCandidate) isPrime() bool {

	sqrtP := int64(math.Floor(math.Sqrt(float64(p))))
	if p%2 == 0 {
		return false
	}
	for divisor := int64(3); divisor <= sqrtP; divisor += 2 {
		if int64(p)%divisor == 0 {
			return false
		}
	}

	return true
}

func (p primeCandidate) isPrimeAKS() bool {

	if p <= 1 {
		return false
	}
	if p <= 3 {
		return true
	}

	if p%2 == 0 || p%3 == 0 {
		return false
	}

	for i := int64(5); i*i <= int64(p); i = i + 6 {
		if int64(p)%i == 0 || int64(p)%(i+2) == 0 {
			return false
		}
	}

	return true

}
