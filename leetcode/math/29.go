package math

func divide(dividend int, divisor int) int {
	abs := func(in int) int {
		if in < 0 {
			in = -in
		}
		return in
	}

	count := 0
	isNegAns := false
	minVal, maxVal := -2147483648, 2147483647

	if dividend/divisor < 0 {
		isNegAns = true
	}
	dividend = abs(dividend)
	divisor = abs(divisor)

	for dividend-divisor >= 0 {
		exp := 0
		for dividend-divisor<<exp >= 0 {
			exp++
		}
		count += 1 << (exp - 1)
		dividend -= divisor << (exp - 1)
	}

	if isNegAns {
		if -count < minVal {
			return minVal
		}
		return -count
	} else {
		if count > maxVal {
			return maxVal
		}
		return count
	}
}
