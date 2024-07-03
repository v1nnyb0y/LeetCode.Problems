package array

import "sort"

func minDifference(nums []int) int {
	if len(nums) <= 4 {
		return 0
	}
	sort.Ints(nums)
	ans := nums[len(nums)-1] - nums[0]
	for i := 0; i <= 3; i++ {
		ans = min(ans, nums[len(nums)-(3-i)-1]-nums[i])
	}
	return ans
}
