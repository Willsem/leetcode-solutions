const mod = 1e9 + 7

func countRoutes(locations []int, start int, finish int, fuel int) int {
	dp := make([][]int, len(locations))
	for i := range dp {
		dp[i] = make([]int, fuel+1)
	}

	for i := 0; i <= fuel; i++ {
		dp[finish][i] = 1
	}

	for j := 0; j <= fuel; j++ {
		for i := 0; i < len(locations); i++ {
			for k := 0; k < len(locations); k++ {
				if k == i {
					continue
				}

				if abs(locations[i]-locations[k]) <= j {
					dp[i][j] = (dp[i][j] + dp[k][j-abs(locations[i]-locations[k])]) % mod
				}
			}
		}
	}

	return dp[start][fuel]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
