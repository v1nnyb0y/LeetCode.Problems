package linked_list

func IsPalindrome(head *ListNode) bool {
	var goForward func(*ListNode) bool
	goForward = func(curr *ListNode) bool {
		if curr.Next != nil {
			result := goForward(curr.Next) && curr.Val == head.Val
			head = head.Next
			return result
		}

		result := head.Val == curr.Val
		head = head.Next
		return result
	}
	curr := head
	return goForward(curr)
}
