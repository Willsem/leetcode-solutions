func change(amount int, coins []int) int {
	n := len(coins)
	dp := make([]int, amount+1)
	dp[0] = 1

	for i := n - 1; i >= 0; i-- {
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j-coins[i]]
		}
	}

	return dp[amount]
}
