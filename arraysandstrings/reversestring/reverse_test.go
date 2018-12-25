package reversestring

import (
	"fmt"
	"golangexercises/arraysandstrings/util"
	"testing"
	"time"
)

const (
	palindrome     = "madamimadam"
	myName         = "My name is a secret and I live in Utah"
	reversedMyName = "hatU ni evil I dna terces a si eman yM"
)

func TestReverseString(t *testing.T) {

	if palindrome != reverse(palindrome) {
		t.Errorf("Error in reversing %s", palindrome)
	}

	if palindrome != improvedReverse(palindrome) {
		t.Errorf("Error in improved reversing %s", palindrome)
	}

	if myName != reverse(reversedMyName) {
		t.Errorf("Error in reversing %s", palindrome)
	}

	if myName != improvedReverse(reversedMyName) {
		t.Errorf("Error in improved reversing %s", palindrome)
	}

	randomString := util.GenerateRandomRune(util.Length)
	start := time.Now()
	reverse(randomString)
	fmt.Printf("Reversed %d elements in %s\n", util.Length, time.Since(start))

	start = time.Now()
	reverse(randomString)
	fmt.Printf("Reversed %d elements in %s using improved method\n", util.Length, time.Since(start))

}
