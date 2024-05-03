package array

func Rotate(matrix [][]int) {
	l, r := 0, len(matrix)-1
	for l < r {
		for i := 0; i < r-l; i++ {
			top, bottom := l, r
			// memo top left
			value := matrix[top][l+i]
			// move bottom left to top left
			matrix[top][l+i] = matrix[bottom-i][l]
			// move bottom right to bottom left
			matrix[bottom-i][l] = matrix[bottom][r-i]
			// move top right to bottom right
			matrix[bottom][r-i] = matrix[top+i][r]
			// move top left to top right
			matrix[top+i][r] = value
		}
		l++
		r--
	}
}
