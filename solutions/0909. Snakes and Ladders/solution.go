func snakesAndLadders(board [][]int) int {
	n := len(board)

	coord := func(k int) (int, int) {
		q, r := (k-1)/n, (k-1)%n
		row := n - 1 - q
		col := r
		if q%2 == 1 {
			col = n - 1 - r
		}
		return row, col
	}

	visited := make([]bool, n*n+1)
	queue := []int{1}
	moves := 0

	for len(queue) > 0 {
		next := []int{}

		for _, curr := range queue {
			if curr == n*n {
				return moves
			}

			for i := 1; i <= 6 && curr+i <= n*n; i++ {
				k := curr + i
				r, c := coord(k)
				if board[r][c] != -1 {
					k = board[r][c]
				}
				if !visited[k] {
					visited[k] = true
					next = append(next, k)
				}
			}
		}

		queue = next
		moves++
	}

	return -1
}
