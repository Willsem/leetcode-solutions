import "math"

func soupServings(n int) float64 {
	m := int(math.Ceil(float64(n) / 25.0))
	dp := make(map[int]map[int]float64, 0)

	for k := 1; k <= m; k++ {
		if calcDp(dp, k, k) > 1-1e-5 {
			return 1.0
		}
	}

	return calcDp(dp, m, m)
}

func calcDp(dp map[int]map[int]float64, i, j int) float64 {
	if i <= 0 && j <= 0 {
		return 0.5
	}
	if i <= 0 {
		return 1.0
	}
	if j <= 0 {
		return 0.0
	}

	if _, ok := dp[i]; !ok {
		dp[i] = make(map[int]float64, 0)
	}

	if _, ok := dp[i][j]; ok {
		return dp[i][j]
	}

	result := (calcDp(dp, i-4, j) +
		calcDp(dp, i-3, j-1) +
		calcDp(dp, i-2, j-2) +
		calcDp(dp, i-1, j-3)) / 4.0

	dp[i][j] = result
	return result
}
