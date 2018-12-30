package hashing

const (
	universeMax    = 1000000
	tableSize      = 100
	numberOfKeys   = 5000
	primeNumber    = 1000003
	hashMultiplier = 51
	hashOffset     = 205
)

func hash(key int) int {
	return ((hashMultiplier*key + hashOffset) % primeNumber) % tableSize
}
