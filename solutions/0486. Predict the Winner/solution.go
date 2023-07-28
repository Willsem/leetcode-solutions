func PredictTheWinner(nums []int) bool {
	n := len(nums)
	dp := make([][]int, n)

	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = nums[i]
	}

	for diff := 1; diff < n; diff++ {
		for left := 0; left < n-diff; left++ {
			right := left + diff
			dp[left][right] = max(
				nums[left]-dp[left+1][right],
				nums[right]-dp[left][right-1],
			)
		}
	}

	return dp[0][n-1] >= 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
