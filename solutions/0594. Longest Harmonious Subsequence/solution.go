func findLHS(nums []int) int {
	counts := make(map[int]int)

	for _, num := range nums {
		counts[num]++
	}

	maxLen := 0
	for num := range counts {
		if _, ok := counts[num-1]; ok {
			maxLen = max(maxLen, counts[num]+counts[num-1])
		}
		if _, ok := counts[num+1]; ok {
			maxLen = max(maxLen, counts[num]+counts[num+1])
		}
	}

	return maxLen
}
