package array

func SetZeroes(matrix [][]int) {
	rows, cols := len(matrix), len(matrix[0])
	rowZero := false

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if matrix[row][col] == 0 {
				matrix[0][col] = 0

				if row > 0 {
					matrix[row][0] = 0
				} else {
					rowZero = true
				}
			}
		}
	}

	for row := 1; row < rows; row++ {
		for col := 1; col < cols; col++ {
			if matrix[0][col] == 0 || matrix[row][0] == 0 {
				matrix[row][col] = 0
			}
		}
	}

	if matrix[0][0] == 0 {
		for row := 0; row < rows; row++ {
			matrix[row][0] = 0
		}
	}

	if rowZero {
		for col := 0; col < cols; col++ {
			matrix[0][col] = 0
		}
	}
}
