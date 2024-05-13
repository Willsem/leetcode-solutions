func matrixScore(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	score := (1 << (n - 1)) * m

	for j := 1; j < n; j++ {
		countSameBits := 0
		for i := 0; i < m; i++ {
			if grid[i][j] == grid[i][0] {
				countSameBits++
			}
		}

		score += (1 << (n - j - 1)) * max(countSameBits, m-countSameBits)
	}

	return score
}
