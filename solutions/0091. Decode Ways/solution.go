import "strconv"

func numDecodings(s string) int {
	dp := make([]int, len(s)+1)
	dp[len(s)] = 1
	for i := len(s) - 1; i > -1; i-- {
		for j := i; j < len(s); j++ {
			n, _ := strconv.Atoi(s[i : j+1])
			if 1 <= n && n <= 26 {
				dp[i] += dp[j+1]
			} else {
				break
			}
		}
	}
	return dp[0]
}
