func maxProfit(prices []int, fee int) int {
	free := 0
	hold := -prices[0]

	for i := 1; i < len(prices); i++ {
		tmp := hold
		hold = max(hold, free-prices[i])
		free = max(free, tmp+prices[i]-fee)
	}

	return free
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
