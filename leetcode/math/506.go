package math

import (
	"sort"
	"strconv"
)

func FindRelativeRanks(score []int) []string {
	var temp []int
	mp := map[int]int{}
	var ans []string

	temp = append(temp, score...)

	sort.Sort(sort.Reverse(sort.IntSlice(temp)))

	for i := range temp {
		mp[temp[i]] = i + 1
	}

	for i := range score {
		if mp[score[i]] == 1 {
			ans = append(ans, "Gold Medal")
		} else if mp[score[i]] == 2 {
			ans = append(ans, "Silver Medal")
		} else if mp[score[i]] == 3 {
			ans = append(ans, "Bronze Medal")
		} else {
			num := strconv.Itoa(mp[score[i]])
			ans = append(ans, num)
		}
	}

	return ans
}
