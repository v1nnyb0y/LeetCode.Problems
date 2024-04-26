package array

func MoveZeroes(nums []int) {
	lastZero := -1
	for i, num := range nums {
		if num != 0 && lastZero != -1 {
			nums[lastZero], nums[i] = nums[i], nums[lastZero]
			if i-lastZero > 1 {
				lastZero++
			} else {
				lastZero = i
			}
		}

		if lastZero == -1 && num == 0 {
			lastZero = i
		}
	}
}
