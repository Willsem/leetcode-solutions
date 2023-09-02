import "math"

func minExtraChar(s string, dictionary []string) int {
	maxVal := len(s) + 1
	dp := make([]int, len(s)+1)
	for i := range dp {
		dp[i] = maxVal
	}
	dp[0] = 0

	dictionarySet := make(map[string]bool)
	for _, word := range dictionary {
		dictionarySet[word] = true
	}

	for i := 1; i <= len(s); i++ {
		dp[i] = dp[i-1] + 1
		for l := 1; l <= i; l++ {
			if dictionarySet[s[i-l:i]] {
				dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-l])))
			}
		}
	}
	return dp[len(s)]
}
