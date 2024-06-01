package string

import "math"

func scoreOfString(s string) int {
	var score int
	// we can use pointer, cause eng symbols only
	for i := 1; i < len(s); i++ {
		score += int(math.Abs(float64(int(s[i-1]) - int(s[i]))))
	}
	return score
}
