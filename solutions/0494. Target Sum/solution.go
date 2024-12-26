import "math"

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum < int(math.Abs(float64(target))) || (sum+target)%2 != 0 {
		return 0
	}

	totalSum := (sum + target) / 2

	dp := make([]int, totalSum+1)
	dp[0] = 1

	for _, num := range nums {
		for i := totalSum; i >= num; i-- {
			dp[i] += dp[i-num]
		}
	}

	return dp[totalSum]
}
