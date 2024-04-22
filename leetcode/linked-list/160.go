package linked_list

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	getLen := func(head *ListNode) int {
		l := 0
		slow := head
		for slow != nil {
			slow = slow.Next
			l++
		}
		return l
	}

	l1 := getLen(headA)
	l2 := getLen(headB)

	a := headA
	b := headB
	if l1 > l2 {
		for a != nil && l1 != l2 {
			a = a.Next
			l1--
		}
	} else if l2 > l1 {
		for b != nil && l1 != l2 {
			b = b.Next
			l2--
		}
	}

	for a != nil && b != nil {
		if a == b {
			return a
		}
		a = a.Next
		b = b.Next
	}
	return nil
}
