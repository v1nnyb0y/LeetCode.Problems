package graph

const (
	LAND  = 1
	WATER = 0
)

type Point struct {
	I, J int
}

func IslandPerimeter(grid [][]int) int {
	var n, m = len(grid), len(grid[0])
	var bfs func(Point) int
	bfs = func(start Point) int {
		var q = []Point{start}
		var dirs = []Point{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
		var visited = make(map[Point]struct{}, 0)
		var perimeter int

		visited[start] = struct{}{}
		for len(q) > 0 {
			curr := q[0]
			q = q[1:]

			for _, dir := range dirs {
				next := Point{curr.I + dir.I, curr.J + dir.J}
				if _, ok := visited[next]; ok {
					continue
				}
				if next.I >= n || next.I < 0 {
					perimeter++
					continue
				}
				if next.J >= m || next.J < 0 {
					perimeter++
					continue
				}
				if grid[next.I][next.J] == WATER {
					perimeter++
					continue
				}
				visited[next] = struct{}{}
				q = append(q, next)
			}
		}
		return perimeter
	}

	for i, row := range grid {
		for j, cell := range row {
			if cell == LAND {
				return bfs(Point{i, j})
			}
		}
	}
	return -1
}
