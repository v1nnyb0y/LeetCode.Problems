package dp

func LongestPalindrome(s string) string {
	dp := make(map[[2]int]bool)
	longestStr := ""
	for i := range s {
		dp[[2]int{i, i}] = true
		longestStr = string(s[i])
	}
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			dp[[2]int{i - 1, i}] = true
			longestStr = s[i-1 : i+1]
		}
	}

	for step := 2; step < len(s); step++ {
		for i := step; i < len(s); i++ {
			isOkByPrev := dp[[2]int{i - step + 1, i - 1}]
			isOkRightNow := s[i] == s[i-step]
			if isOkByPrev && isOkRightNow {
				dp[[2]int{i - step, i}] = true
				longestStr = s[i-step : i+1]
			}
		}
	}
	return longestStr
}
