const mod = 1e9 + 7

func numTilings(n int) int {
	dp := make([]int, n+3)
	dp[0], dp[1], dp[2] = 1, 2, 5

	for i := 3; i < n; i++ {
		dp[i] = (2*dp[i-1] + dp[i-3]) % mod
	}

	return dp[n-1]
}
