const (
	POINTS     = 0
	BRAINPOWER = 1
)

func mostPoints(questions [][]int) int64 {
	dp := make([]int64, len(questions)+1)

	for i := len(questions) - 1; i >= 0; i-- {
		dp[i] = int64(questions[i][POINTS])

		nextQuestion := i + questions[i][BRAINPOWER] + 1
		if nextQuestion < len(questions) {
			dp[i] += dp[nextQuestion]
		}

		dp[i] = max(dp[i], dp[i+1])
	}

	return dp[0]
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}

	return b
}
