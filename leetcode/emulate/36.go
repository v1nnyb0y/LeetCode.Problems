package emulate

func IsValidSudoku(board [][]byte) bool {
	var rows, cols, square [9][9]bool

	for i, row := range board {
		for j, v := range row {
			if v == '.' {
				continue
			}
			unit := int(v) - 48
			if rows[i][unit-1] || cols[j][unit-1] || square[i/3*3+j/3][unit-1] {
				return false
			}
			rows[i][unit-1] = true
			cols[j][unit-1] = true
			square[i/3*3+j/3][unit-1] = true
		}
	}

	return true
}
