package string

import (
	"strings"
	"unicode"
)

func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	left, right := 0, len(s)-1

	isValidRune := func(symb rune) bool {
		return unicode.IsLetter(symb) || unicode.IsDigit(symb)
	}

	for left < right {
		validLeft, validRight := rune(s[left]), rune(s[right])
		if !isValidRune(validLeft) {
			left++
			continue
		}
		if !isValidRune(validRight) {
			right--
			continue
		}
		if validLeft != validRight {
			return false
		}
		left++
		right--
	}
	return true
}
