package binarytrees

import "fmt"

type Node struct {
	Value                 int
	LeftChild, RightChild *Node
}

func New(v int) *Node {
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
