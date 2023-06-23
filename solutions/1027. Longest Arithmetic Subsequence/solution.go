func longestArithSeqLength(nums []int) int {
	maxLen := 0
	dp := make([]map[int]int, len(nums))

	for right := 0; right < len(nums); right++ {
		dp[right] = make(map[int]int, 0)
		for left := 0; left < right; left++ {
			diff := nums[left] - nums[right]
			if _, ok := dp[left][diff]; !ok {
				dp[right][diff] = 2
			} else {
				dp[right][diff] = dp[left][diff] + 1
			}

			if maxLen < dp[right][diff] {
				maxLen = dp[right][diff]
			}
		}
	}

	return maxLen
}
