package trees

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Left != nil {
		invertTree(root.Left)
	}
	if root.Right != nil {
		invertTree(root.Right)
	}

	root.Left, root.Right = root.Right, root.Left

	return root
}
