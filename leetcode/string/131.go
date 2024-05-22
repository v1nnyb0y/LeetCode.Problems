package string

func partition(s string) [][]string {
	isPalindrome := func(s string) bool {
		left, right := 0, len(s)-1
		for left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
		return true
	}

	var result [][]string
	var curr []string
	lenS := len(s)

	var explore func(index int)
	explore = func(index int) {
		if index >= lenS {
			result = append(result, append([]string(nil), curr...))
			return
		}

		for i := index; i < lenS; i++ {
			subStr := s[index : i+1]
			if isPalindrome(subStr) {
				curr = append(curr, subStr)
				explore(i + 1)
				curr = curr[:len(curr)-1]
			}
		}
	}

	explore(0)
	return result
}
