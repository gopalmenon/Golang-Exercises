package checkifbalanced

import (
	"golangexercises/binarytrees"
	"golangexercises/binarytrees/fillfromarray"
	"testing"
)

func TestBalanced(t *testing.T) {

	intArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	bstRoot1 := fillfromarray.CreateTree(intArray)
	if !IsBalanced(bstRoot1) {
		t.Errorf("Expected balanced tree\n")
	}

	bstRoot2 := binarytrees.New(1)
	bstRoot2.RightChild = binarytrees.New(2)
	bstRoot2.RightChild.RightChild = binarytrees.New(3)

	if IsBalanced(bstRoot2) {
		t.Errorf("Expected unbalanced tree\n")
	}

}
