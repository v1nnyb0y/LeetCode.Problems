package linked_list

type listNode struct {
	Val  int
	Next *listNode
}

func MergeTwoLists(list1 *listNode, list2 *listNode) *listNode {
	newList := listNode{-1, nil}
	head := &newList
	for list1 != nil || list2 != nil {
		val1, val2 := 101, 101
		if list1 != nil {
			val1 = list1.Val
		}

		if list2 != nil {
			val2 = list2.Val
		}

		if val1 < val2 {
			head.Next = &listNode{val1, nil}
			list1 = list1.Next
		} else {
			head.Next = &listNode{val2, nil}
			list2 = list2.Next
		}
		head = head.Next
	}

	return newList.Next
}
