package dynamicprogramming

import (
	"fmt"
	"testing"
)

var k1 = []rune("malayalam")
var k2 = []rune("pneumonoultramicroscopicsilicovolcanoconiosis")

func TestLPSS(t *testing.T) {

	if lpssLengthK1, lpssK1 := LongestPalindromicSubSequence(k1, 0, len(k1)-1); lpssLengthK1 != len(k1) {
		t.Errorf("LPSS of %v should be %d long.\n", k1, len(k1))
	} else {
		fmt.Printf("LPSS in %s is %s and is %d long\n", string(k1), string(lpssK1), lpssLengthK1)
	}

	cache = nil

	lpssLengthK2, lpssK2 := LongestPalindromicSubSequence(k2, 0, len(k2)-1)
	fmt.Printf("LPSS in %s is %s and is %d long\n", string(k2), string(lpssK2), lpssLengthK2)

}
