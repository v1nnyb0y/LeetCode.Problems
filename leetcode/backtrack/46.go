package backtrack

func Permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{[]int{nums[0]}}
	}

	var result [][]int
	i := 0
	for i < len(nums) {
		value := nums[0]
		perms := Permute(nums[1:])
		nums = append(nums[1:], value)

		for _, perm := range perms {
			perm = append(perm, value)
			result = append(result, perm)
		}
		i++
	}
	return result
}
