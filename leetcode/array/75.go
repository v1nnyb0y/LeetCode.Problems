package array

func SortColors(nums []int) {
	var counts [3]int
	for _, num := range nums {
		counts[num]++
	}

	var j int
	for color, count := range counts {
		var i int
		for i != count {
			nums[j] = color
			j++
			i++
		}
	}
}
