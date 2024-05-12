package array

func LargestLocal(grid [][]int) [][]int {
	I, J := len(grid)-2, len(grid[0])-2
	for i, row := range grid[:I] {
		row[0] = max(row[0], grid[i+1][0], grid[i+2][0])
		row[1] = max(row[1], grid[i+1][1], grid[i+2][1])
		for j, x := range row[:J] {
			row[j+2] = max(row[j+2], grid[i+1][j+2], grid[i+2][j+2])
			row[j] = max(x, row[j+1], row[j+2])
		}
		grid[i] = row[:J]
	}
	return grid[:I]
}
