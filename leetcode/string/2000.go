package string

func ReversePrefix(word string, ch byte) string {
	var start, left int
	left = -1
	for i := 0; i < len(word); i++ {
		if word[i] == ch {
			left = i
			break
		}
	}

	if left == -1 {
		return word
	}
	wordArr := []rune(word)
	for start < left {
		wordArr[start], wordArr[left] = wordArr[left], wordArr[start]
		start++
		left--
	}
	return string(wordArr)
}
