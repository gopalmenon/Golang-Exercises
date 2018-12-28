package stacksandqueues

type Node struct {
	Value interface{}
	next  *Node
}

type Stack struct {
	top *Node
}

func New() *Stack {
	s := new(Stack)
	return s
}

func (s *Stack) Pop() interface{} {

	if s.top == nil {
		return nil
	}

	retval := s.top.Value
	s.top = s.top.next
	return retval

}

func (s *Stack) Push(v interface{}) {

	n := &Node{Value: v}
	n.next = s.top
	s.top = n

}

func (s *Stack) Peek() interface{} {
	return s.top.Value
}
