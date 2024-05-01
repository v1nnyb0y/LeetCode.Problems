package emulate

func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	phone := make(map[rune]string)
	phone['2'] = "abc"
	phone['3'] = "def"
	phone['4'] = "ghi"
	phone['5'] = "jkl"
	phone['6'] = "mno"
	phone['7'] = "pqrs"
	phone['8'] = "tuv"
	phone['9'] = "wxyz"

	var comb []string
	var backtrack func(d, result string)
	backtrack = func(d, result string) {
		if len(result) == len(digits) {
			comb = append(comb, result)
			return
		}

		for _, digit := range phone[rune(d[0])] {
			result += string(digit)
			backtrack(d[1:], result)
			result = result[:len(result)-1]
		}
	}

	backtrack(digits, "")
	return comb
}
