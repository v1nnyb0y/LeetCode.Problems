package string

func LongestCommonPrefix(strs []string) string {
	var commonPrefix string = strs[0]
	for _, str := range strs[1:] {
		var s1, s2 int
		for s1 < len(commonPrefix) &&
			s2 < len(str) &&
			commonPrefix[s1] == str[s2] {
			s1++
			s2++
		}

		commonPrefix = commonPrefix[:s1]
		if len(commonPrefix) == 0 {
			break
		}
	}
	return commonPrefix
}
