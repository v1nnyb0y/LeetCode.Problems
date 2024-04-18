package string

func check(hs string, ndl string) bool {
	var i, j = 0, 0
	for i < len(hs) && j < len(ndl) && hs[i] == ndl[j] {
		i++
		j++
	}
	return j == len(ndl)
}

func StrStr(hs string, ndl string) int {
	for ind, chr := range hs {
		if chr == rune(ndl[0]) && check(hs[ind:], ndl) {
			return ind
		}
	}
	return -1
}
