package linked_list

func DoubleIt(head *ListNode) *ListNode {
	reverse := func(node *ListNode) *ListNode {
		var cur, prev *ListNode
		cur, prev = node, nil
		for cur != nil {
			next := cur.Next
			cur.Next = prev
			cur, prev = next, cur
		}
		return prev
	}

	back := reverse(head)
	node := back
	var memo int
	var prev *ListNode
	for node != nil {
		val := node.Val*2 + memo
		if val > 9 {
			node.Val = val - 10
			memo = 1
		} else {
			node.Val = val
			memo = 0
		}
		prev, node = node, node.Next
	}

	if memo == 1 {
		prev.Next = &ListNode{Val: 1}
	}
	return reverse(back)
}
