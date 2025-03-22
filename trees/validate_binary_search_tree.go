package trees

func isValidBST(root *TreeNode) bool {
	var prev *TreeNode
	return inorder(root, &prev)
}

func inorder(root *TreeNode, prev **TreeNode) bool {
	if root == nil {
		return true
	}

	if !inorder(root.Left, prev) {
		return false
	}

	if *prev != nil && root.Val <= (*prev).Val {
		return false
	}
	*prev = root

	return inorder(root.Right, prev)
}
