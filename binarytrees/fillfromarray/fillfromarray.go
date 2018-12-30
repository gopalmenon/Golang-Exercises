package fillfromarray

import (
	"golangexercises/binarytrees"
)

//If array is sorted, it returns a BST root node
func createTree(array []int) *binarytrees.Node {

	if len(array) == 0 {
		return nil
	}

	n := binarytrees.New(array[len(array)/2])
	n.LeftChild = createTree(array[:len(array)/2])
	n.RightChild = createTree(array[len(array)/2+1:])
	return n

}
