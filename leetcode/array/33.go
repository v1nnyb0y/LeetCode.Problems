package array

func Search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (r + l) / 2
		left, mid, right := nums[l], nums[m], nums[r]

		if mid == target {
			return m
		}
		if left <= mid {
			if target >= left && target < mid {
				r = m - 1
			} else {
				l = m + 1
			}
		} else {
			if target > mid && target <= right {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}
	return -1
}
