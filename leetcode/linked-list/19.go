package linked_list

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	p := head
	listSize := 0
	for p != nil {
		p = p.Next
		listSize++
	}

	if listSize == n {
		return head.Next
	}
	listSize -= n

	curr := 0
	p = head
	for curr != listSize-1 {
		p = p.Next
		curr++
	}
	p.Next = p.Next.Next
	return head
}
