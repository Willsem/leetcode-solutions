func timeRequiredToBuy(tickets []int, k int) int {
	ans := 0
	for i := range tickets {
		if i <= k {
			ans += min(tickets[k], tickets[i])
		} else {
			ans += min(tickets[k]-1, tickets[i])
		}
	}
	return ans
}
