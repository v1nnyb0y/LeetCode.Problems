package linked_list

func RemoveNodes(head *ListNode) *ListNode {
	ptr := head
	arr := []int{}
	for ptr != nil {
		arr = append(arr, ptr.Val)
		ptr = ptr.Next
	}
	mx := -1
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] >= mx {
			mx = arr[i]
		} else {
			arr[i] = -1
		}
	}
	dum := &ListNode{}
	ap := dum
	cnt := 0
	p := head
	for p != nil {
		if arr[cnt] != -1 {
			ap.Next = p
			ap = ap.Next
		}
		cnt++
		p = p.Next
	}
	return dum.Next
}
