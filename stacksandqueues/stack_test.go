package stacksandqueues

import (
	"fmt"
	"strings"
	"testing"
)

func TestStack(t *testing.T) {

	s := New()
	s.push(9)
	s.push(0)
	s.push(0)
	s.push(4)
	s.push(8)

	if v := s.peek(); v != 8 {
		t.Errorf("Peeked in stack. Expected 8 but got %d\n", v)
	}

	builder := strings.Builder{}
	for v := s.pop(); v != nil; v = s.pop() {
		builder.WriteString(fmt.Sprintf("%d", v))
	}

	if builder.String() != "84009" {
		t.Errorf("Expected 84009 but got %s\n", builder.String())
	}

}
