package array

import "sort"

func MaximumHappinessSum(hs []int, k int) int64 {
	sort.Sort(sort.Reverse(sort.IntSlice(hs)))
	var hsum int64
	for i, h := range hs[:k] {
		val := int64(h - i)
		if val > 0 {
			hsum += val
		}
	}
	return hsum
}
