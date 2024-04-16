package tree

func AddOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if root == nil {
		return nil
	}

	if depth == 1 {
		root = &TreeNode{Val: val, Left: root}
	} else if depth == 2 {
		root.Left = &TreeNode{Val: val, Left: root.Left}
		root.Right = &TreeNode{Val: val, Right: root.Right}
	} else {
		AddOneRow(root.Left, val, depth-1)
		AddOneRow(root.Right, val, depth-1)
	}

	return root
}
