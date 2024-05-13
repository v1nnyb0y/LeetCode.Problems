package graph

func Exist(board [][]byte, word string) bool {
	var dfs func(int, int, int) bool
	dfs = func(i, j, count int) bool {
		if count == len(word) {
			return true
		}

		if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != word[count] {
			return false
		}

		temp := board[i][j]
		board[i][j] = '*'
		found := dfs(i+1, j, count+1) ||
			dfs(i-1, j, count+1) ||
			dfs(i, j+1, count+1) ||
			dfs(i, j-1, count+1)
		board[i][j] = temp
		return found
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] && dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
