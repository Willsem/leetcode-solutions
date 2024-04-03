func exist(board [][]byte, word string) bool {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == word[0] && dfs(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, word string, i, j, curr int) bool {
	if curr == len(word) {
		return true
	}

	if i < 0 || i >= len(board) || j < 0 || j >= len(board[i]) || board[i][j] != word[curr] {
		return false
	}

	tmp := board[i][j]
	defer func() {
		board[i][j] = tmp
	}()

	board[i][j] = '*'
	return dfs(board, word, i+1, j, curr+1) ||
		dfs(board, word, i-1, j, curr+1) ||
		dfs(board, word, i, j+1, curr+1) ||
		dfs(board, word, i, j-1, curr+1)
}
