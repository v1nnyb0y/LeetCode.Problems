package dp

func CanJump(nums []int) bool {
	dp := make([]bool, len(nums))
	dp[0] = true
	i := 0
	for i < len(nums) {
		if dp[i] {
			for j := i; j < len(nums) && j <= i+nums[i]; j++ {
				dp[j] = true
			}
		}
		i++
	}

	return dp[len(nums)-1]
}
