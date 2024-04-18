package math

func PlusOne(digits []int) []int {
	var memo = 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += memo
		if digits[i] >= 10 {
			digits[i] %= 10
			memo = 1
		} else {
			memo = 0
		}
	}

	if memo == 1 {
		result := []int{1}
		result = append(result, digits...)
		return result
	}
	return digits
}
