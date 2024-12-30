const mod = 1e9 + 7

func countGoodStrings(low, high, zero, one int) int {
	dp := make([]int, high+1)
	dp[0] = 1
	ans := 0

	for i := 1; i <= high; i++ {
		if i >= zero {
			dp[i] += dp[i-zero]
		}

		if i >= one {
			dp[i] += dp[i-one]
		}

		dp[i] %= mod

		if i >= low && i <= high {
			ans = (ans + dp[i]) % mod
		}
	}

	return ans
}
