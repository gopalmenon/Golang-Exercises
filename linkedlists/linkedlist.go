package linkedlists

import (
	"fmt"
	"strings"
)

type Node struct {
	list  *List
	Next  *Node
	Value interface{}
}

func (n *Node) NextNode() *Node {
	return n.Next
}

type List struct {
	Head *Node
}

func (l *List) init() {
	l.Head = nil
}

func New() *List {
	l := new(List)
	l.init()
	return l
}

func (l *List) Pushback(v interface{}) *Node {

	newNode := &Node{list: l, Value: v}
	if l.Head == nil {
		l.Head = newNode
	} else {
		node, nextNode := l.Head, l.Head.Next
		for ; nextNode != nil; node, nextNode = nextNode, nextNode.Next {
		}
		node.Next = newNode
	}

	return newNode
}

func (l *List) Remove(n *Node) interface{} {

	if n.list == l {
		node, nextNode := l.Head, l.Head.Next
		if node == n {
			l.Head = node.Next
			return n.Value
		}
		for ; nextNode != nil; node, nextNode = nextNode, nextNode.Next {
			if nextNode == n {
				node.Next = nextNode.Next
			}
		}
	}

	return n.Value
}

func (l *List) RemoveValue(v interface{}) bool {

	for node := l.Head; node != nil; node = node.Next {
		if node.Value == v {
			l.Remove(node)
			return true
		}
	}

	return false
}

func (l *List) String() string {
	var str strings.Builder
	for node := l.Head; node != nil; node = node.Next {
		str.WriteString(fmt.Sprint(node.Value))
		str.WriteString(",")
	}
	return str.String()
}
