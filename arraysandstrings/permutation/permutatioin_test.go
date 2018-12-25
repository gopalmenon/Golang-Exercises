package permutation

import (
	"fmt"
	"golangexercises/arraysandstrings/util"
	"testing"
	"time"
)

func TestPermutationString(t *testing.T) {

	if !IsPermutationOfOtherUsingSort(util.MyName, util.ReversedMyName) {
		t.Errorf("%s not found to be permutation of %s using sort\n", util.MyName, util.ReversedMyName)
	}

	if !IsPermutationOfOtherUsingHashing(util.MyName, util.ReversedMyName) {
		t.Errorf("%s not found to be permutation of %s using hashing\n", util.MyName, util.ReversedMyName)
	}

	randomString1, randomString2 := util.GenerateRandomRune(util.Length), util.GenerateRandomRune(util.Length)
	start := time.Now()
	IsPermutationOfOtherUsingSort(randomString1, randomString2)
	fmt.Printf("Checked permutation of %d elements in %s using sort\n", util.Length, time.Since(start))

	start = time.Now()
	IsPermutationOfOtherUsingHashing(randomString1, randomString2)
	fmt.Printf("Checked permutation of %d elements in %s using hashing\n", util.Length, time.Since(start))

}
