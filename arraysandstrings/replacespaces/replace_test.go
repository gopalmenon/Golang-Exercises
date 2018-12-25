package replacespaces

import (
	"fmt"
	"golangexercises/arraysandstrings/util"
	"strings"
	"testing"
	"time"
)

const (
	myName         = "My name is a secret and I live in Utah                  "
	myNameReplaced = "My%20name%20is%20a%20secret%20and%20I%20live%20in%20Utah"
)

func TestReplace(t *testing.T) {

	if r := ReplaceSpaces(myName); r != myNameReplaced {
		t.Errorf("%s is not correct replacement for %s\n", r, myName)
	}

	randomString := util.GenerateRandomRune(util.Length)
	start := time.Now()
	ReplaceSpaces(randomString)
	fmt.Printf("Replaced spaces in %d elements in %s\n", util.Length, time.Since(start))

	start = time.Now()
	r := strings.NewReplacer(" ", "%20")
	r.Replace(randomString)
	fmt.Printf("Replaced spaces in %d elements in %s using Go strings Replacer\n", util.Length, time.Since(start))

}
