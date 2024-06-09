package array

func subarraysDivByK(nums []int, k int) int {
	remainders := make(map[int]int)
	remainders[0] = 1

	currentSum := 0
	cnt := 0

	for i := 0; i < len(nums); i++ {
		currentSum += nums[i]
		remainder := currentSum % k

		if remainder < 0 {
			remainder += k
		}

		cnt += remainders[remainder]
		remainders[remainder]++
	}

	return cnt
}
