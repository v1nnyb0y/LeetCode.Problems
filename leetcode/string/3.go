package string

func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	left := 0
	longest := -1 << 31
	for left < len(s) {
		memo := make(map[byte]int)
		for left < len(s) {
			if prev, ok := memo[s[left]]; ok {
				left = prev + 1
				break
			}
			memo[s[left]] = left
			left++
		}
		longest = max(longest, len(memo))
	}
	return longest
}
