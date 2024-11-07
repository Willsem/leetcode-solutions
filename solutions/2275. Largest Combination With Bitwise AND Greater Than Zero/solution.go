func largestCombination(candidates []int) int {
	maxCount := 0
	for i := 0; i < 24; i++ {
		count := 0
		for _, num := range candidates {
			if num&(1<<i) != 0 {
				count++
			}
		}
		maxCount = max(maxCount, count)
	}
	return maxCount
}
