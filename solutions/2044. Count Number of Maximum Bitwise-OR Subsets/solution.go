func countMaxOrSubsets(nums []int) int {
	maxValue := 0
	dp := make([]int, 1<<17)
	dp[0] = 1

	for _, num := range nums {
		for i := maxValue; i >= 0; i-- {
			dp[i|num] += dp[i]
		}
		maxValue |= num
	}

	return dp[maxValue]
}
