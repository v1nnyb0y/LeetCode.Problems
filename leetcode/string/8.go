package string

import (
	"math"
	"unicode"
)

func MyAtoi(s string) int {
	sign := 1
	result := 0
	i := 0
	countSigns := 0
	for i < len(s) && !unicode.IsDigit(rune(s[i])) {
		char := string(s[i])

		if char == "+" {
			if countSigns == 1 {
				return 0
			}
			sign = 1
			countSigns++
		} else if char == "-" {
			if countSigns == 1 {
				return 0
			}
			sign = -1
			countSigns++
		} else if char != " " {
			return 0
		} else if countSigns == 1 {
			return 0
		}
		i++
	}

	for i < len(s) && unicode.IsDigit(rune(s[i])) {
		char := int(s[i] - '0')
		result *= 10
		result += char
		i++

		if result > math.MaxInt32 {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}
	}

	result *= sign
	return result
}
