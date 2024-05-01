package math

import "sort"

func ThreeSum(nums []int) [][]int {
	visited := make(map[[3]int]bool)
	sums := make(map[int][]int)
	for i, num := range nums {
		sums[num] = append(sums[num], i)
	}

	for i, val := range nums {

		for j := i + 1; j+1 < len(nums); j++ {
			val2 := nums[j]
			k, ok := sums[0-val-val2]
			if ok && k[len(k)-1] != i && k[len(k)-1] != j {
				key := []int{val, val2, nums[k[len(k)-1]]}
				sort.Ints(key)
				if _, ok2 := visited[[3]int(key)]; !ok2 {
					visited[[3]int(key)] = true
				}
			}
		}
	}

	var result [][]int
	for value := range visited {
		result = append(result, []int{value[0], value[1], value[2]})
	}
	return result
}
