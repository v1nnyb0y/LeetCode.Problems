package array

func MaxSubArray(nums []int) int {
	maxSub := nums[0]
	currSum := 0

	for _, value := range nums {
		if currSum < 0 {
			currSum = 0
		}
		currSum += value
		maxSub = max(maxSub, currSum)
	}
	return maxSub
}
