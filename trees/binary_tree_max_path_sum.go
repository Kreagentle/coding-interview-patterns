package trees

import "math"

func maxPathSum(root *TreeNode) int {
	_, res := helperMaxPathSum(root, math.MinInt32)
	return res
}

func helperMaxPathSum(root *TreeNode, res int) (int, int) {
	if root == nil {
		return 0, 0
	}

	var leftChild, rightChild int
	if root.Left != nil {
		leftChild, res = helperMaxPathSum(root.Left, res)
		leftChild = max(leftChild, 0)
	}

	if root.Right != nil {
		rightChild, res = helperMaxPathSum(root.Right, res)
		rightChild = max(rightChild, 0)
	}

	newPath := root.Val + leftChild + rightChild

	res = max(res, newPath)
	maxContribution := root.Val + max(leftChild, rightChild)
	return maxContribution, res
}
