func minCostClimbingStairs(cost []int) int {
	if len(cost) == 1 {
		return cost[0]
	}

	for i := 2; i < len(cost); i++ {
		cost[i] = min(cost[i]+cost[i-1], cost[i]+cost[i-2])
	}

	return min(cost[len(cost)-1], cost[len(cost)-2])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
