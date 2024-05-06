package math

import "math"

func MyPow(x float64, n int) float64 {
	var helper func(float64, int) float64
	helper = func(x float64, n int) float64 {
		if x == 0 {
			return 0
		}
		if n == 0 {
			return 1
		}

		var res = helper(x, n/2)
		res = res * res

		if n%2 != 0 {
			res = res * x
		}

		return res
	}

	var res = helper(x, int(math.Abs(float64(n))))

	if n < 0 {
		return 1 / res
	}
	return res
}
