package linked_list

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	incrementIfNotNull := func(node **ListNode, sum *int) {
		if *node != nil {
			*sum += (*node).Val
			*node = (*node).Next
		}
	}
	createNodeWithSum := func(node **ListNode, sum int) {
		(*node).Next = &ListNode{Val: sum}
		*node = (*node).Next
	}

	head := &ListNode{}
	headP := head

	memo := 0
	for l1 != nil || l2 != nil {
		var sum = memo
		incrementIfNotNull(&l1, &sum)
		incrementIfNotNull(&l2, &sum)

		if sum >= 10 {
			memo = 1
			sum %= 10
		} else {
			memo = 0
		}

		createNodeWithSum(&headP, sum)
	}

	if memo != 0 {
		createNodeWithSum(&headP, 1)
	}
	return head.Next
}
