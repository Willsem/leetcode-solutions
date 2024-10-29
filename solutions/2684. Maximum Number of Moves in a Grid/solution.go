func maxMoves(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := range dp {
		dp[i] = make([]int, len(grid[i]))
		dp[i][0] = 1
	}

	maxMoves := 0
	for j := 1; j < len(grid[0]); j++ {
		for i := 0; i < len(grid); i++ {
			if i > 0 && grid[i-1][j-1] < grid[i][j] && dp[i-1][j-1] > 0 {
				dp[i][j] = dp[i-1][j-1] + 1
			}

			if grid[i][j-1] < grid[i][j] && dp[i][j-1] > 0 {
				dp[i][j] = max(dp[i][j], dp[i][j-1]+1)
			}

			if i < len(grid)-1 && grid[i+1][j-1] < grid[i][j] && dp[i+1][j-1] > 0 {
				dp[i][j] = max(dp[i][j], dp[i+1][j-1]+1)
			}

			maxMoves = max(maxMoves, dp[i][j])
		}
	}

	if maxMoves == 0 {
		return 0
	}
	return maxMoves - 1
}
