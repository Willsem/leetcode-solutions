const mod = 1e9 + 7

func numMusicPlaylists(n int, goal int, k int) int {
	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 1

	for i := 1; i <= goal; i++ {
		dp[i%2][0] = 0
		for j := 1; j <= min(i, n); j++ {
			dp[i%2][j] = dp[(i-1)%2][j-1] * (n - (j - 1)) % mod
			if j > k {
				dp[i%2][j] = (dp[i%2][j] + dp[(i-1)%2][j]*(j-k)) % mod
			}
		}
	}

	return dp[goal%2][n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
