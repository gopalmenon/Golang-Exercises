package checkifbalanced

import "golangexercises/binarytrees"

func TreeDepth(treeRootNode *binarytrees.Node) (int, int) {

	if treeRootNode == nil {
		return 0, 0
	}

	leftDepthMax, leftDepthMin := TreeDepth(treeRootNode.LeftChild)
	leftDepthMax += 1
	leftDepthMin += 1
	rightDepthMax, rightDepthMin := TreeDepth(treeRootNode.RightChild)
	rightDepthMax += 1
	rightDepthMin += 1

	var maxDepth, minDepth int
	if leftDepthMax > rightDepthMax {
		maxDepth = leftDepthMax
	} else {
		maxDepth = rightDepthMax
	}
	if leftDepthMin < rightDepthMin {
		minDepth = leftDepthMin
	} else {
		minDepth = rightDepthMin
	}

	return maxDepth, minDepth
}

func IsBalanced(treeRootNode *binarytrees.Node) bool {

	maxDepth, minDepth := TreeDepth(treeRootNode)
	if maxDepth-minDepth > 1 {
		return false
	}

	return true

}
