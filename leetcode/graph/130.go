package graph

func solve(board [][]byte) {
	var WATER byte = 'X'
	rows, cols := len(board), len(board[0])
	dirs := [][2]int{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}
	visited := make(map[[2]int]struct{})

	var findIgnore func(int, int)
	findIgnore = func(r, c int) {
		if r < 0 || r >= rows {
			return
		}
		if c < 0 || c >= cols {
			return
		}
		if _, ok := visited[[2]int{r, c}]; ok {
			return
		}
		if board[r][c] == WATER {
			return
		}

		visited[[2]int{r, c}] = struct{}{}
		for _, dir := range dirs {
			findIgnore(r+dir[0], c+dir[1])
		}
	}

	for r := 0; r < rows; r++ {
		findIgnore(r, 0)
		findIgnore(r, cols-1)
	}

	for c := 0; c < cols; c++ {
		findIgnore(0, c)
		findIgnore(rows-1, c)
	}

	for r := range board {
		for c := range board[r] {
			if _, ok := visited[[2]int{r, c}]; !ok {
				board[r][c] = WATER
			}
		}
	}
}
