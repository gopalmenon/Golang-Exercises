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

func (s *Stack) pop() interface{} {

	if s.top == nil {
		return nil
	}

	retval := s.top.Value
	s.top = s.top.next
	return retval

}

func (s *Stack) push(v interface{}) {

	n := &Node{Value: v}
	n.next = s.top
	s.top = n

}

func (s *Stack) peek() interface{} {
	return s.top.Value
}
