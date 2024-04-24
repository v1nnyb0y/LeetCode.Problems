package dp

import (
	"fmt"
	"strconv"
)

func IsHappy(n int) bool {
	memo := make(map[int]struct{})
	for {
		if _, ok := memo[n]; ok {
			break
		}
		memo[n] = struct{}{}

		sum := 0
		for _, cd := range strconv.Itoa(n) {
			d, _ := strconv.Atoi(string(cd))
			sum += d * d
		}
		fmt.Println(sum)
		if sum == 1 {
			return true
		}
		n = sum
	}
	return false
}
