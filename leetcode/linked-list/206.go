package linked_list

func ReverseList(head *ListNode) *ListNode {
	var curr, prev *ListNode
	curr = head
	prev = nil

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
