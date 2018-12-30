package binarytrees

import "fmt"

type Node struct {
	Value                 interface{}
	LeftChild, RightChild *Node
}

func New(v interface{}) *Node {
	return &Node{Value: v}
}

func InOrderTraversal(rootNode *Node) {

	if rootNode == nil {
		return
	}

	InOrderTraversal(rootNode.LeftChild)
	fmt.Printf("%s\t", rootNode.Value)
	InOrderTraversal(rootNode.RightChild)

}
