package string

func TitleToNumber(columnTitle string) int {

	multiplier := 1
	result := 0

	for i := len(columnTitle) - 1; i >= 0; i -= 1 {
		result += (int(columnTitle[i]) - int('A') + 1) * multiplier
		multiplier *= 26
	}

	return result
}
