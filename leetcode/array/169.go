package array

func MajorityElement(nums []int) int {
	mCount := 1
	mVal := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == mVal {
			mCount++
			continue
		}

		mCount--
		if mCount == 0 {
			mVal = nums[i]
			mCount = 1
		}
	}
	return mVal
}
