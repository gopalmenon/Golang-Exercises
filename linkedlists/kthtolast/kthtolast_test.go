package kthtolast

import (
	"fmt"
	"golangexercises/linkedlists"
	"math/rand"
	"testing"
)

func TestKThToLast(t *testing.T) {

	l := linkedlists.New()
	for counter := 0; counter < 100; counter++ {
		l.Pushback(rand.Intn(100))
	}

	k := rand.Intn(100)
	kToLastElement := kthToLastElement(l, k)
	fmt.Printf("%d to last element of %s is %d\n", k, l, kToLastElement.Value)

}
