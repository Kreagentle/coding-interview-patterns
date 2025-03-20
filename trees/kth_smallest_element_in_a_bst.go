package trees

func kthSmallest(root *TreeNode, k int) int {
	var res int
	helper(root, &k, &res)
	return res
}

func helper(root *TreeNode, k *int, res *int) {
	if root == nil {
		return
	}

	if root.Left != nil {
		helper(root.Left, k, res)
	}

	*k--
	if *k == 0 {
		*res = root.Val
		return
	}

	if root.Right != nil {
		helper(root.Right, k, res)
	}

	return
}
