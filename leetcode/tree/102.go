package tree

func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	queue := [][]*TreeNode{[]*TreeNode{root}}
	for len(queue) > 0 {
		resultRow := []int{}

		newRow := []*TreeNode{}
		prevRow := queue[0]

		queue = queue[1:]

		for _, node := range prevRow {
			resultRow = append(resultRow, node.Val)
			if node.Left != nil {
				newRow = append(newRow, node.Left)
			}
			if node.Right != nil {
				newRow = append(newRow, node.Right)
			}
		}

		result = append(result, resultRow)
		if len(newRow) > 0 {
			queue = append(queue, newRow)
		}
	}

	return result
}
