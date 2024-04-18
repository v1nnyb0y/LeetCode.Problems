package math

func MySqrt(x int) int {
	left, right := 0, x/2+2

	for left < right {
		mid := left + (right-left)/2
		if mid*mid > x {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left - 1
}
