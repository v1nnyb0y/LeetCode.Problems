package tree

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := MaxDepth(root.Left), MaxDepth(root.Right)
	return max(left, right) + 1
}
