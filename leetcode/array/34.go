package array

func SearchRange(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	for l <= r {
		index := l + (r-l)/2
		mid := nums[index]
		if mid == target {
			r = index
			for ; r < len(nums)-1; r++ {
				if nums[r+1] != target {
					break
				}
			}
			l = index
			for ; l >= 1; l-- {
				if nums[l-1] != target {
					break
				}
			}

			return []int{l, r}
		} else if mid > target {
			r--
		} else {
			l++
		}
	}

	return []int{-1, -1}
}
