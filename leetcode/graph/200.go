package graph

func NumIslands(grid [][]byte) int {
	count := 0

	bfs := func(grid [][]byte, r, c int) {
		q := [][2]int{}

		q = append(q, [2]int{r, c})
		grid[r][c] = '2'

		offsets := [4][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
		for len(q) > 0 {
			curr := q[0]
			q = q[1:]

			for _, offset := range offsets {
				x := curr[0] + offset[0]
				y := curr[1] + offset[1]

				if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[x]) && grid[x][y] == '1' {
					q = append(q, [2]int{x, y})
					grid[x][y] = 2
				}
			}
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				count++

				bfs(grid, i, j)
			}
		}
	}

	return count
}
