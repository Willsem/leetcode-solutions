func numberOfArithmeticSlices(nums []int) int {
	n, res := len(nums), 0

	dp := make([]map[int]int, n)
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			diff := nums[i] - nums[j]

			if dp[i] == nil {
				dp[i] = make(map[int]int)
			}

			dp[i][diff]++

			if _, ok := dp[j][diff]; ok {
				dp[i][diff] += dp[j][diff]
				res += dp[j][diff]
			}
		}
	}

	return res
}
