package tree

func SortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	left, right := 0, len(nums)-1
	mid := (left + right) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = SortedArrayToBST(nums[:mid])
	root.Right = SortedArrayToBST(nums[mid+1:])
	return root
}
