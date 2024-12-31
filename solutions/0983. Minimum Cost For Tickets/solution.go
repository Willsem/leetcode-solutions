func mincostTickets(days []int, costs []int) int {
	lastDay := days[len(days)-1]
	dp := make([]int, lastDay+1)

	for d, i := 1, 0; d <= lastDay; d++ {
		if d < days[i] {
			dp[d] = dp[d-1]
		} else {
			i++

			cost1 := dp[d-1] + costs[0]
			cost2 := dp[max(0, d-7)] + costs[1]
			cost3 := dp[max(0, d-30)] + costs[2]
			dp[d] = min(cost1, cost2, cost3)
		}
	}

	return dp[lastDay]
}
