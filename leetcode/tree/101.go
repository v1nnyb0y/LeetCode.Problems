package tree

func IsSymmetric(root *TreeNode) bool {

	var dfs func(*TreeNode, *TreeNode) bool
	dfs = func(left *TreeNode, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil {
			return false
		}

		return left.Val == right.Val && dfs(left.Left, right.Right) && dfs(left.Right, right.Left)
	}
	return dfs(root.Left, root.Right)
}
