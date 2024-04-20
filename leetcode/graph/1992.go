package graph

const (
	FOREST = 0
)

func FindFarmland(land [][]int) [][]int {
	n, m := len(land), len(land[0])

	dirs := [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	var dfs func(int, int) (int, int)
	dfs = func(i, j int) (ei, ej int) {
		ei, ej = i, j
		land[i][j] = FOREST

		for _, dir := range dirs {
			newI, newJ := i+dir[0], j+dir[1]
			if newI >= n ||
				newI < 0 ||
				newJ >= m ||
				newJ < 0 ||
				land[newI][newJ] == FOREST {
				continue
			}
			cei, cej := dfs(newI, newJ)
			ei, ej = max(ei, cei), max(ej, cej)
		}
		return
	}

	result := [][]int{}
	for i, row := range land {
		for j, cell := range row {
			if cell == LAND {
				ei, ej := dfs(i, j)
				result = append(result, []int{i, j, ei, ej})
			}
		}
	}
	return result
}
