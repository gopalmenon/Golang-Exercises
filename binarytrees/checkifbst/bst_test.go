package checkifbst

import (
	"golangexercises/binarytrees/fillfromarray"
	"sort"
	"testing"
)

func TestBST(t *testing.T) {

	intArray := []int{2, 15, 3, 4, 5, 11, 6, 7, 9, 8, 10, 12, 1, 13, 14}
	bstRoot1 := fillfromarray.CreateTree(intArray)
	if IsBST(bstRoot1) {
		t.Errorf("Not a BST\n")
	}

	sort.Ints(intArray)
	bstRoot2 := fillfromarray.CreateTree(intArray)
	if !IsBST(bstRoot2) {
		t.Errorf("Is a BST\n")
	}
}
