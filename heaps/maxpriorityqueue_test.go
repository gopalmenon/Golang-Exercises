package heaps

import (
	"fmt"
	"testing"
)

func TestMaxPriorityQueue(t *testing.T) {

	array := []int{9, 7, 8, 5, 6, 4, 3, 2, 0, 1}
	mpq := New(array)
	if maxPeek, ok := mpq.Max(); ok {
		if maxPeek != 9 {
			t.Errorf("Expected 9 but got %d\n", maxPeek)
		}
	} else {
		t.Errorf("Expected 9 but got error\n")
	}

	if max, ok := mpq.ExtractMax(); ok {
		if max != 9 {
			t.Errorf("Expected 9 but got %d\n", max)
		}
	} else {
		t.Errorf("Expected 9 but got error\n")
	}

	mpq.Insert(100)
	if maxPeek, ok := mpq.Max(); ok {
		if maxPeek != 100 {
			t.Errorf("Expected 100 but got %d\n", maxPeek)
		}
	} else {
		t.Errorf("Expected 100 but got error\n")
	}

	fmt.Printf("Max Heap:\n")
	String(mpq.array)

}
