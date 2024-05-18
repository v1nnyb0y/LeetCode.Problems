package tree

func distributeCoins(root *TreeNode) int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	var moves int

	var traverse func(node *TreeNode) int
	traverse = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftExcess := traverse(node.Left)
		rightExcess := traverse(node.Right)
		totalExcess := node.Val + leftExcess + rightExcess - 1
		moves += abs(totalExcess)
		return totalExcess
	}

	traverse(root)
	return moves
}
