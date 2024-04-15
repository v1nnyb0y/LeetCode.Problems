package tree

type treeNode struct {
	Val   int
	Left  *treeNode
	Right *treeNode
}

func dfs(root *treeNode, prevSum int) int {
	var newSum = prevSum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return newSum
	}

	var result int
	if root.Left != nil {
		result += dfs(root.Left, newSum)
	}
	if root.Right != nil {
		result += dfs(root.Right, newSum)
	}
	return result
}

func SumNumbers(root *treeNode) int {
	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	var sum int
	if root.Left != nil {
		sum += dfs(root.Left, root.Val)
	}
	if root.Right != nil {
		sum += dfs(root.Right, root.Val)
	}
	return sum
}
