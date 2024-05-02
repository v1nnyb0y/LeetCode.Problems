package array

import "math"

func FindMaxK(nums []int) int {
	memo := make(map[int]struct{})
	maxEl := math.MinInt32

	for _, num := range nums {
		if num < 0 {
			if _, ok := memo[-num]; ok {
				maxEl = max(maxEl, -num)
			} else {
				memo[num] = struct{}{}
			}
		} else {
			if _, ok := memo[-num]; ok {
				maxEl = max(maxEl, num)
			} else {
				memo[num] = struct{}{}
			}
		}
	}

	if maxEl == math.MinInt32 {
		return -1
	}
	return maxEl
}
