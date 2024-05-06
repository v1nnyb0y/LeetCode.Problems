package array

import "sort"

func merge(intervals [][]int) [][]int {
	var result [][]int

	sort.Slice(
		intervals, func(i, j int) bool {
			return intervals[i][0] < intervals[j][0]
		},
	)

	startCurr, endCurr := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		start, end := intervals[i][0], intervals[i][1]
		if endCurr >= start {
			startCurr = min(startCurr, start)
			endCurr = max(endCurr, end)
		} else {
			result = append(result, []int{startCurr, endCurr})
			startCurr, endCurr = start, end
		}
	}
	result = append(result, []int{startCurr, endCurr})
	return result
}
