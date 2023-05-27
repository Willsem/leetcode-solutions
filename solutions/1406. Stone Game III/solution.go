func stoneGameIII(stoneValue []int) string {
	dp := make([]int, 4)
	k, s := 0, 0
	for i := len(stoneValue) - 1; i >= 0; i-- {
		k = (k + 1) % 4
		s = stoneValue[i]
		// take only one stone
		dp[k] = s - dp[(k+3)%4]
		// let's try take two and three stones
		for j := 1; j < 3 && i+j < len(stoneValue); j++ {
			s += stoneValue[i+j]
			if r := s - dp[(k+3-j)%4]; r > dp[k] {
				dp[k] = r
			}
		}
	}
	switch {
	case dp[k] > 0:
		return "Alice"
	case dp[k] < 0:
		return "Bob"
	}
	return "Tie"
}
