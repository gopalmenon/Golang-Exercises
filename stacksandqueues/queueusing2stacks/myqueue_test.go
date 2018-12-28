package queueusing2stacks

import (
	"fmt"
	"strings"
	"testing"
)

func TestQueue(t *testing.T) {

	q := New()
	q.Enqueue(8)
	q.Enqueue(4)

	v := q.Dequeue()

	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("%d", v))

	q.Enqueue(0)
	q.Enqueue(0)
	q.Enqueue(9)

	for v := q.Dequeue(); v != nil; v = q.Dequeue() {
		builder.WriteString(fmt.Sprintf("%d", v))
	}
	if builder.String() != "84009" {
		t.Errorf("Expected 84009 but got %s\n", builder.String())
	}

}
