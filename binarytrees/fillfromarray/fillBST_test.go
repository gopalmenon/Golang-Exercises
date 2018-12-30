package fillfromarray

import (
	"golangexercises/binarytrees"
	"testing"
)

func TestFillBST(t *testing.T) {

	intArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	bstRoot := CreateTree(intArray)
	binarytrees.InOrderTraversal(bstRoot)

}
