func maxSumAfterPartitioning(arr []int, k int) int {
	dp := make([]int, len(arr)+1)

	for i := 1; i <= len(arr); i++ {
		maximum, maxSum := 0, 0
		for j := 1; j <= k && i-j >= 0; j++ {
			maximum = max(maximum, arr[i-j])
			maxSum = max(maxSum, dp[i-j]+maximum*j)
		}
		dp[i] = maxSum
	}

	return dp[len(arr)]
}
