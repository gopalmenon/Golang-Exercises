package linkedlists

import (
	"fmt"
	"testing"
)

func TestLink(t *testing.T) {

	l := New()
	elements := make([]*Node, 10)

	for i := 0; i < 10; i++ {
		elements[i] = l.Pushback(i)
	}

	l.Remove(elements[7])

	fmt.Println(l)
	if l.String() != "0,1,2,3,4,5,6,8,9," {
		t.Errorf("List should contain %s instead of %s\n", "012345689", l.String())
	}

	l.Remove(elements[0])
	l.Remove(elements[1])
	l.Remove(elements[2])
	l.Remove(elements[3])

	if l.RemoveValue(4) == false {
		t.Errorf("4 should have been found to remove\n")
	}

	l.Remove(elements[5])
	l.Remove(elements[6])
	l.Remove(elements[8])

	fmt.Println(l)
	if l.String() != "9," {
		t.Errorf("List should contain %s instead of %s\n", "9", l.String())
	}
	l.Remove(elements[9])

	fmt.Println(l)
	if l.String() != "" {
		t.Errorf("List should be empty instead of %s\n", l.String())
	}
}
