package string

func ReverseString(s []byte) {
	p1, p2 := 0, len(s)-1
	for p1 < p2 {
		s[p1], s[p2] = s[p2], s[p1]
		p1 += 1
		p2 -= 1
	}
}
