package reversestring

import (
	"fmt"
	"golangexercises/arraysandstrings/util"
	"testing"
	"time"
)

func TestReverseString(t *testing.T) {

	if util.Palindrome != reverse(util.Palindrome) {
		t.Errorf("Error in reversing %s\n", util.Palindrome)
	}

	if util.Palindrome != improvedReverse(util.Palindrome) {
		t.Errorf("Error in improved reversing %s\n", util.Palindrome)
	}

	if util.MyName != reverse(util.ReversedMyName) {
		t.Errorf("Error in reversing %s\n", util.Palindrome)
	}

	if util.MyName != improvedReverse(util.ReversedMyName) {
		t.Errorf("Error in improved reversing %s\n", util.Palindrome)
	}

	randomString := util.GenerateRandomRune(util.Length)
	start := time.Now()
	reverse(randomString)
	fmt.Printf("Reversed %d elements in %s\n", util.Length, time.Since(start))

	start = time.Now()
	reverse(randomString)
	fmt.Printf("Reversed %d elements in %s using improved method\n", util.Length, time.Since(start))

}
