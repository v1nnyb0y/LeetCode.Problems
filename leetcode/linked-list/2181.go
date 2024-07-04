package linked_list

func mergeNodes(head *ListNode) *ListNode {
	newHead := &ListNode{}
	pointer := newHead

	head = head.Next

	var sum int
	for head != nil {
		val := head.Val
		if val == 0 {
			node := &ListNode{Val: sum}
			pointer.Next = node
			pointer = pointer.Next
			sum = 0
		} else {
			sum += val
		}
		head = head.Next
	}

	return newHead.Next
}
