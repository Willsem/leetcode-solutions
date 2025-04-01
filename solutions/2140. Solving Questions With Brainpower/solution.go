func mostPoints(questions [][]int) int64 {
	dp := make([]int, len(questions)+1)

	for i := len(questions) - 1; i >= 0; i-- {
		points, brainpower := questions[i][0], questions[i][1]

		dp[i] = points

		next := i + brainpower + 1
		if next < len(questions) {
			dp[i] += dp[next]
		}

		dp[i] = max(dp[i], dp[i+1])
	}

	return int64(dp[0])
}
