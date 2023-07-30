func strangePrinter(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var solve func(left, right int) int
	solve = func(left, right int) int {
		if dp[left][right] != -1 {
			return dp[left][right]
		}

		dp[left][right] = n
		j := -1

		for i := left; i < right; i++ {
			if s[i] != s[right] && j == -1 {
				j = i
			}

			if j != -1 {
				dp[left][right] = min(dp[left][right], solve(j, i)+solve(i+1, right)+1)
			}
		}

		if j == -1 {
			dp[left][right] = 0
		}

		return dp[left][right]
	}

	return solve(0, n-1) + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
