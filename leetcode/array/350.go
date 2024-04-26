package array

import "sort"

func Intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	p1, p2 := 0, 0
	var res []int
	for p1 < len(nums1) && p2 < len(nums2) {
		if nums1[p1] == nums2[p2] {
			res = append(res, nums1[p1])
			p1 += 1
			p2 += 1
			continue
		}
		if nums1[p1] > nums2[p2] {
			p2 += 1
		} else {
			p1 += 1
		}
	}
	return res
}
