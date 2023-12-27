func minCost(colors string, neededTime []int) int {
	var cx int32
	var ci int
	var totalCost int

	for i, c := range colors {
		if cx != c {
			cx = c
			ci = i
		} else {
			costCx := neededTime[ci]
			costC := neededTime[i]
			if costC > costCx {
				totalCost += costCx
				ci = i
			} else {
				totalCost += costC
			}
		}
	}

	return totalCost
}
