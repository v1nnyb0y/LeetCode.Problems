package bit

func MissingNumber(nums []int) int {
	n := len(nums)
	res := n
	for i := 0; i < n; i++ {
		res = res ^ i ^ nums[i]
	}
	// after the loop we have
	// XOR sum of all the numbers twice + one number once
	return res
}
