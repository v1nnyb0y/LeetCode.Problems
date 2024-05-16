package tree

func EvaluateTree(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return root.Val == 1
	}

	leftBool := false
	if root.Left != nil && root.Left.Val > 1 {
		leftBool = EvaluateTree(root.Left)
	} else {
		leftBool = root.Left.Val == 1
	}

	rightBool := false
	if root.Right != nil && root.Right.Val > 1 {
		rightBool = EvaluateTree(root.Right)
	} else {
		rightBool = root.Right.Val == 1
	}

	if root.Val == 2 {
		return leftBool || rightBool
	} else {
		return leftBool && rightBool
	}
}
