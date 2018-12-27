package removeduplicates

import (
	"fmt"
	"golangexercises/linkedlists"
	"math/rand"
	"testing"
	"time"
)

func TestDedup(t *testing.T) {

	dedupFuncs := []func(l *linkedlists.List){removeDupsWithHashing, removeDupsWithRunners}
	for fNumber, f := range dedupFuncs {

		l := linkedlists.New()

		l.Pushback(1)
		l.Pushback(2)
		l.Pushback(2)
		l.Pushback(3)
		l.Pushback(3)
		l.Pushback(3)
		l.Pushback(4)
		l.Pushback(4)
		l.Pushback(4)
		l.Pushback(4)
		l.Pushback(5)
		l.Pushback(5)
		l.Pushback(5)
		l.Pushback(5)
		l.Pushback(5)
		f(l)
		if "1,2,3,4,5," != l.String() {
			t.Errorf("List should contain %s instead of %s\n", "1,2,3,4,5,", l.String())
		}

		l = linkedlists.New()
		for counter := 0; counter < 100; counter++ {
			l.Pushback(rand.Intn(100))
		}

		fmt.Printf("Remove duplicates from: %s\n", l)
		start := time.Now()
		f(l)
		fmt.Printf("Removed duplicated using function %d in %s\n", fNumber, time.Since(start))
		fmt.Printf("After removing duplicates: %s\n", l)
	}
}
