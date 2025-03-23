package trees

func verticalOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	columns := make(map[int][]int)
	queue := []Obj{{node: root, col: 0}}
	columns[0] = []int{root.Val}

	minCol, maxCol := 0, 0

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		if item.node.Left != nil {
			queue = append(queue, Obj{node: item.node.Left, col: item.col - 1})
			columns[item.col-1] = append(columns[item.col-1], item.node.Left.Val)
			minCol = min(minCol, item.col-1)
		}

		if item.node.Right != nil {
			queue = append(queue, Obj{node: item.node.Right, col: item.col + 1})
			columns[item.col+1] = append(columns[item.col+1], item.node.Right.Val)
			maxCol = max(maxCol, item.col+1)
		}
	}

	result := make([][]int, maxCol-minCol+1)
	for i := minCol; i <= maxCol; i++ {
		result[i-minCol] = columns[i]
	}

	return result
}

type Obj struct {
	col  int
	node *TreeNode
}
