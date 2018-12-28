package stacksandqueues

type Queue struct {
	head *Node
	tail *Node
}

func NewQueue() *Queue {
	return new(Queue)
}

func (q *Queue) Enqueue(v interface{}) {

	n := &Node{Value: v}
	if q.head == nil {
		q.tail = n
		q.head = q.tail
	} else {
		q.tail.next = n
		q.tail = q.tail.next
	}

}

func (q *Queue) Dequeue() interface{} {

	if q.head == nil {
		return nil
	}

	retval := q.head.Value
	q.head = q.head.next
	return retval

}
