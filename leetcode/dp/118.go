package dp

func Generate(numRows int) [][]int {
	result := [][]int{{1}, {1, 1}}
	if numRows == 1 {
		return result[:1]
	}

	numRows -= 2
	for numRows != 0 {
		row := []int{1}
		prev := result[len(result)-1]
		for i := 1; i < len(prev); i++ {
			el := prev[i-1] + prev[i]
			row = append(row, el)
		}
		row = append(row, 1)
		result = append(result, row)
		numRows--
	}
	return result
}
