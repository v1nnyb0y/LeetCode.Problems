package array

func TwoSum(nums []int, target int) []int {
	memo := make(map[int]int, len(nums)-1)
	for id, num := range nums {
		if fId, ok := memo[target-num]; ok {
			return []int{fId, id}
		}
		memo[num] = id
	}
	return []int{-1, -1}
}
