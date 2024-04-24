package dp

func Tribonacci(n int) int {
	t0, t1, t2 := 0, 1, 1
	if n == 0 {
		return t0
	}
	if n == 1 {
		return t1
	}

	for n != 2 {
		value := t0 + t1 + t2
		t0, t1, t2 = t1, t2, value
		n--
	}
	return t2
}
