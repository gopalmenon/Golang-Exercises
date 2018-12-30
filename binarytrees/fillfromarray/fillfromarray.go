package fillfromarray

import (
	"golangexercises/binarytrees"
)

//If array is sorted, it returns a BST root node
func CreateTree(array []int) *binarytrees.Node {

	if len(array) == 0 {
		return nil
	}

	n := binarytrees.New(array[len(array)/2])
	n.LeftChild = CreateTree(array[:len(array)/2])
	n.RightChild = CreateTree(array[len(array)/2+1:])
	return n

}
