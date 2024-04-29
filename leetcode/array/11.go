package array

func MaxArea(h []int) int {
	l, r := 0, len(h)-1
	maxS := 0
	for l < r {
		height := min(h[l], h[r])
		width := r - l
		maxS = max(maxS, height*width)

		if h[l] > h[r] {
			r--
		} else {
			l++
		}
	}
	return maxS
}
