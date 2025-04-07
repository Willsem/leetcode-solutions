func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%2 != 0 {
		return false
	}
	targetSum := sum / 2

	dp := make([]bool, targetSum+1)
	dp[0] = true

	for _, num := range nums {
		for i := targetSum; i >= num; i-- {
			dp[i] = dp[i] || dp[i-num]
		}
	}

	return dp[targetSum]
}
