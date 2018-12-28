package stacksandqueues

import (
	"fmt"
	"strings"
	"testing"
)

func TestStack(t *testing.T) {

	s := New()
	s.Push(9)
	s.Push(0)
	s.Push(0)
	s.Push(4)
	s.Push(8)

	if v := s.Peek(); v != 8 {
		t.Errorf("Peeked in stack. Expected 8 but got %d\n", v)
	}

	builder := strings.Builder{}
	for v := s.Pop(); v != nil; v = s.Pop() {
		builder.WriteString(fmt.Sprintf("%d", v))
	}

	if builder.String() != "84009" {
		t.Errorf("Expected 84009 but got %s\n", builder.String())
	}

}
