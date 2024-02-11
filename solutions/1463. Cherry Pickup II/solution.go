func cherryPickup(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])

	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, m)
		for j := range dp[i] {
			dp[i][j] = make([]int, m)
		}
	}

	cherries := 0
	dp[0][0][m-1] = grid[0][0] + grid[0][m-1]
	for i := 1; i < n; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < m; k++ {
				if j > i || k < m-i-1 || j > k {
					continue
				}
				dp[i][j][k] = dp[i-1][j][k]
				if j > 0 {
					dp[i][j][k] = max(dp[i][j][k], dp[i-1][j-1][k])
				}
				if j > 0 && k > 0 {
					dp[i][j][k] = max(dp[i][j][k], dp[i-1][j-1][k-1])
				}
				if j > 0 && k+1 < m {
					dp[i][j][k] = max(dp[i][j][k], dp[i-1][j-1][k+1])
				}
				if j+1 < m {
					dp[i][j][k] = max(dp[i][j][k], dp[i-1][j+1][k])
				}
				if j+1 < m && k > 0 {
					dp[i][j][k] = max(dp[i][j][k], dp[i-1][j+1][k-1])
				}
				if j+1 < m && k+1 < m {
					dp[i][j][k] = max(dp[i][j][k], dp[i-1][j+1][k+1])
				}
				if k > 0 {
					dp[i][j][k] = max(dp[i][j][k], dp[i-1][j][k-1])
				}
				if k+1 < m {
					dp[i][j][k] = max(dp[i][j][k], dp[i-1][j][k+1])
				}
				if j != k {
					dp[i][j][k] += grid[i][j] + grid[i][k]
				} else {
					dp[i][j][k] += grid[i][j]
				}
				cherries = max(cherries, dp[i][j][k])
			}
		}
	}

	return cherries
}
