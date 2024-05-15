package graph

func maximumSafenessFactor(grid [][]int) int {
	var (
		colors = make([]int, len(grid)*len(grid))       // visited cell marks
		sCache = make(map[int]bool)                     // safeness cache, to make sure we don't calculate the possibility more than once
		queue  = make([][2]int, 0, len(grid)*len(grid)) // BSF queue
		qT, qH = 0, 0                                   // BFS queue pointers
	)

	// Find thieves
	for i, row := range grid {
		for j, cell := range row {
			if cell == 1 {
				// append as BFS starting points
				queue = append(queue, [2]int{i, j})
				qT++
			}
		}
	}

	// this BFS calculates distances from thieves to each cell (+1)
	for qH < qT {
		cell := queue[qH]
		qH++
		if cell[0] > 0 && (grid[cell[0]-1][cell[1]] == 0 || grid[cell[0]-1][cell[1]] > grid[cell[0]][cell[1]]+1) {
			grid[cell[0]-1][cell[1]] = grid[cell[0]][cell[1]] + 1
			queue = append(queue, [2]int{cell[0] - 1, cell[1]})
			qT++
		}
		if cell[0] < len(grid)-1 && (grid[cell[0]+1][cell[1]] == 0 || grid[cell[0]+1][cell[1]] > grid[cell[0]][cell[1]]+1) {
			grid[cell[0]+1][cell[1]] = grid[cell[0]][cell[1]] + 1
			queue = append(queue, [2]int{cell[0] + 1, cell[1]})
			qT++
		}
		if cell[1] > 0 && (grid[cell[0]][cell[1]-1] == 0 || grid[cell[0]][cell[1]-1] > grid[cell[0]][cell[1]]+1) {
			grid[cell[0]][cell[1]-1] = grid[cell[0]][cell[1]] + 1
			queue = append(queue, [2]int{cell[0], cell[1] - 1})
			qT++
		}
		if cell[1] < len(grid)-1 && (grid[cell[0]][cell[1]+1] == 0 || grid[cell[0]][cell[1]+1] > grid[cell[0]][cell[1]]+1) {
			grid[cell[0]][cell[1]+1] = grid[cell[0]][cell[1]] + 1
			queue = append(queue, [2]int{cell[0], cell[1] + 1})
			qT++
		}
	}

	// this BFS finds path between top-left to bottom-right
	// with specified minimum safeness
	var pathWithMinSafeness = func(safeness int) (exists bool) {
		// I suspect that bin search loop may query pathWithMinSafeness
		// more than once for the same safeness, so I cache previously calculated result.
		if v, ok := sCache[safeness]; ok {
			return v
		}
		if grid[0][0] < safeness {
			sCache[safeness] = false
			return
		}

		// clear queue
		queue = queue[:0]
		queue = append(queue, [2]int{0, 0})
		qH, qT = 0, 1

		for qH < qT {
			cell := queue[qH] // take element from q head
			qH++
			if cell[0] == len(grid)-1 && cell[1] == len(grid)-1 {
				// reached end, specified safeness is posible
				exists = true
				break
			}
			// I use 'colors' here to check is specified cell was visied
			// when searching path for particular safeness.
			// The trick with using 'safeness' as color value is
			// to avoid clearing or reallocating colors before each BFS run.

			// Check if can move to neighbouring cells, considering borders, safeness and colors.
			if cell[0] < len(grid)-1 && grid[cell[0]+1][cell[1]] >= safeness && colors[(cell[0]+1)*len(grid)+cell[1]] != safeness {
				colors[(cell[0]+1)*len(grid)+cell[1]] = safeness
				queue = append(queue, [2]int{cell[0] + 1, cell[1]})
				qT++
			}
			if cell[1] < len(grid)-1 && grid[cell[0]][cell[1]+1] >= safeness && colors[cell[0]*len(grid)+cell[1]+1] != safeness {
				colors[cell[0]*len(grid)+cell[1]+1] = safeness
				queue = append(queue, [2]int{cell[0], cell[1] + 1})
				qT++
			}
			if cell[0] > 0 && grid[cell[0]-1][cell[1]] >= safeness && colors[(cell[0]-1)*len(grid)+cell[1]] != safeness {
				colors[(cell[0]-1)*len(grid)+cell[1]] = safeness
				queue = append(queue, [2]int{cell[0] - 1, cell[1]})
				qT++
			}
			if cell[1] > 0 && grid[cell[0]][cell[1]-1] >= safeness && colors[cell[0]*len(grid)+cell[1]-1] != safeness {
				colors[cell[0]*len(grid)+cell[1]-1] = safeness
				queue = append(queue, [2]int{cell[0], cell[1] - 1})
				qT++
			}
		}
		sCache[safeness] = exists // cache result
		return
	}

	// Find maximum safeness with bin search
	var (
		l, r     = 2, grid[0][0] + 1 // range is [l; r)
		safeness = 1
	)
	for l < r {
		mid := (l + r) / 2
		if !pathWithMinSafeness(mid) {
			// no path for this safeness
			// goto left half
			r = mid
		} else {
			// have path for this safeness
			if !pathWithMinSafeness(mid + 1) {
				// +1 safeness is not possible,
				// mid is the answer
				safeness = mid
				break
			}
			// current and +1 safenesses are possible,
			// check right half
			l = mid + 1
		}
	}
	// return ans-1, because inital distances were calculated as +1
	return safeness - 1
}
