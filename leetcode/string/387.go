package string

func FirstUniqChar(s string) int {
	dict := make(map[rune][]int)
	for i, chr := range s {
		dict[chr] = append(dict[chr], i)
	}
	firstOqq := 100001
	for _, oqq := range dict {
		if len(oqq) == 1 && firstOqq > oqq[0] {
			firstOqq = oqq[0]
		}
	}
	if firstOqq == 100001 {
		return -1
	}
	return firstOqq
}
