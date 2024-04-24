package array

func ContainsDuplicate(nums []int) bool {
	memo := make(map[int]struct{})
	for _, val := range nums {
		if _, ok := memo[val]; ok {
			return true
		}
		memo[val] = struct{}{}
	}
	return false
}
