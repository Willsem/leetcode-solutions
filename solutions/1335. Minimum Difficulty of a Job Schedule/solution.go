import "math"

func minDifficulty(jobDifficulty []int, d int) int {
	n := len(jobDifficulty)
	if n < d {
		return -1
	}

	dp := make([]int, n)
	dp[0] = jobDifficulty[0]
	for i := 1; i < n; i++ {
		dp[i] = int(math.Max(float64(dp[i-1]), float64(jobDifficulty[i])))
	}

	for i := 1; i < d; i++ {
		for j := n - 1; j >= i; j-- {
			dp[j] = math.MaxInt
			mx := 0
			for k := j; k >= i; k-- {
				mx = int(math.Max(float64(mx), float64(jobDifficulty[k])))
				dp[j] = int(math.Min(float64(dp[j]), float64(dp[k-1]+mx)))
			}
		}
	}

	return dp[n-1]
}
