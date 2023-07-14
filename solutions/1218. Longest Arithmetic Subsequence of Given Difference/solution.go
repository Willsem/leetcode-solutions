func longestSubsequence(arr []int, difference int) int {
	dp := make(map[int]int, 0)
	answer := 1
	for _, v := range arr {
		if _, ok := dp[v-difference]; !ok {
			dp[v-difference] = 0
		}

		dp[v] = dp[v-difference] + 1

		if answer < dp[v] {
			answer = dp[v]
		}
	}

	return answer
}
