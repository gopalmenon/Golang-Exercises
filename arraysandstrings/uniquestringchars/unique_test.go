package uniquestringchars

import (
	"fmt"
	"golangexercises/arraysandstrings/util"
	"testing"
	"time"
)

func TestUniqueChars(t *testing.T) {

	testStringDups := "Madam I'm Adam"
	testStringUnique := "This man"

	if HasUniqueCharsBySorting(testStringDups) {
		t.Errorf("Could not locate duplicates in %s\n", testStringDups)
	}

	if HasUniqueCharsByHashing(testStringDups) {
		t.Errorf("Could not locate duplicates in %s\n", testStringDups)
	}

	if !HasUniqueCharsBySorting(testStringUnique) {
		t.Errorf("Found duplicates in %s\n", testStringUnique)
	}

	if !HasUniqueCharsByHashing(testStringUnique) {
		t.Errorf("Found duplicates in %s\n", testStringUnique)
	}

	randomString := util.GenerateRandomRune(util.Length)
	start := time.Now()
	HasUniqueCharsBySorting(randomString)
	fmt.Printf("Checked %d elements in %s using sorting\n", util.Length, time.Since(start))
	start = time.Now()
	HasUniqueCharsByHashing(randomString)
	fmt.Printf("Checked %d elements in %s using hashing\n", util.Length, time.Since(start))
	start = time.Now()

}
