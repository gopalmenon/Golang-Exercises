package hashing

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestUniversalHashing(t *testing.T) {

	m := make(map[int][]int, numberOfKeys)

	for counter := 0; counter < numberOfKeys; counter++ {
		randomKey := rand.Intn(universeMax)
		hashValue := hash(randomKey)
		m[hashValue] = append(m[hashValue], randomKey)
	}

	fmt.Println("Hash\t\tCount\t\tKeys")
	for hashValue, keys := range m {
		fmt.Printf("%d\t\t\t%d\t\t\t%v\n", hashValue, len(keys), keys)
	}

	//Show ascii histogram
	for counter := 0; counter < tableSize; counter++ {
		fmt.Printf("%d\t\t", counter)
		for xCounter := 0; xCounter < len(m[counter]); xCounter++ {
			fmt.Printf("x")
		}
		fmt.Printf("\n")
	}

}
