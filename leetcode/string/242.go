package string

func IsAnagram(s string, t string) bool {
	dict := make(map[rune]int)
	for _, chr := range s {
		dict[chr]++
	}
	for _, chr := range t {
		dict[chr]--
	}
	for _, cnt := range dict {
		if cnt != 0 {
			return false
		}
	}
	return true
}
