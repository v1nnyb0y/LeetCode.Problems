package backtrack

func Subsets(nums []int) [][]int {
	var backtrack func(*[][]int, []int, []int, int)
	backtrack = func(result *[][]int, current []int, nums []int, start int) {
		subsetCopy := make([]int, len(current))
		copy(subsetCopy, current)
		*result = append(*result, subsetCopy)

		for i := start; i < len(nums); i++ {
			current = append(current, nums[i])
			backtrack(result, current, nums, i+1)
			current = current[:len(current)-1]
		}
	}

	var result [][]int
	backtrack(&result, []int{}, nums, 0)
	return result
}
