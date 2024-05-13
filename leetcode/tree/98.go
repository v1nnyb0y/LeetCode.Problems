package tree

import "math"

func isValidBST(root *TreeNode) bool {
	return valid(root, math.MinInt64, math.MaxInt64)
}

func valid(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	return valid(root.Left, min, root.Val) && valid(root.Right, root.Val, max)
}
