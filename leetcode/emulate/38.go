package emulate

import "strconv"

func CountAndSay(n int) string {
	unionSay := func(str string) string {
		result := ""
		i := 0
		for i < len(str) {
			val := str[i]
			j := i
			for j < len(str) && val == str[j] {
				j++
			}
			result += strconv.Itoa(j-i) + string(val)
			i = j
		}

		return result
	}

	if n == 1 {
		return "1"
	}
	result := CountAndSay(n - 1)
	result = unionSay(result)
	return result
}
