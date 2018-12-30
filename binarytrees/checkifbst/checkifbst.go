package checkifbst

import "golangexercises/binarytrees"

func IsBST(rootNode *binarytrees.Node) bool {

	if rootNode == nil {
		return true
	}

	childrenBSTs := IsBST(rootNode.LeftChild) && IsBST(rootNode.RightChild)

	selfBST := true

	if (rootNode.LeftChild != nil && (rootNode.LeftChild.Value > rootNode.Value)) ||
		(rootNode.RightChild != nil && (rootNode.RightChild.Value <= rootNode.Value)) {
		selfBST = false
	}

	return childrenBSTs && selfBST

}
