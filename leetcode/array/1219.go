package array

func getMaximumGold(grid [][]int) int {

	m := len(grid)
	n := len(grid[0])

	var dfs func(grid [][]int, i, j, sum int) int
	dfs = func(grid [][]int, i, j, sum int) int {
		if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] == 0 {
			return sum
		}

		currentGold := grid[i][j]
		sum += currentGold

		grid[i][j] = 0

		left := dfs(grid, i, j-1, sum)
		right := dfs(grid, i, j+1, sum)
		top := dfs(grid, i-1, j, sum)
		bottom := dfs(grid, i+1, j, sum)

		grid[i][j] = currentGold

		return max(left, right, top, bottom)
	}

	highest := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res := dfs(grid, i, j, 0)

			if res > highest {
				highest = res
			}
		}
	}

	return highest
}
