package uniquestringchars

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const length = 10000

func TestUniqueChars(t *testing.T) {

	testStringDups := "Madam I'm Adam"
	testStringUnique := "This man"

	if HasUniqueCharsBySorting(testStringDups) {
		t.Errorf("Could not locate duplicates in %s", testStringDups)
	}

	if HasUniqueCharsByHashing(testStringDups) {
		t.Errorf("Could not locate duplicates in %s", testStringDups)
	}

	if !HasUniqueCharsBySorting(testStringUnique) {
		t.Errorf("Found duplicates in %s", testStringUnique)
	}

	if !HasUniqueCharsByHashing(testStringUnique) {
		t.Errorf("Found duplicates in %s", testStringUnique)
	}

	randomString := generateRandomRune(length)
	start := time.Now()
	HasUniqueCharsBySorting(randomString)
	fmt.Printf("Checked %d elements in %s using sorting\n", length, time.Since(start))
	start = time.Now()
	HasUniqueCharsByHashing(randomString)
	fmt.Printf("Checked %d elements in %s using hashing\n", length, time.Since(start))
	start = time.Now()

}

var runes = []rune("一二三四五六七八九十1234567890")

func generateRandomRune(n int) string {
	randRune := make([]rune, n)

	for i := range randRune {
		// without this, the final value will be same all the time.
		rand.Seed(time.Now().UnixNano())

		randRune[i] = runes[rand.Intn(len(runes))]
	}
	return string(randRune)
}
