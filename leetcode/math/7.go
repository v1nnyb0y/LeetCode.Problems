package math

import "math"

func Reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
	}

	x *= sign
	res := 0
	for x != 0 {
		res *= 10
		res += x % 10
		x /= 10

		if res*sign >= math.MaxInt32 || res*sign <= math.MinInt32 {
			return 0
		}
	}

	return res * sign
}
