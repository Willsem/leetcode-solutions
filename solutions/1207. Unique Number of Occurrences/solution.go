func uniqueOccurrences(arr []int) bool {
	counts := make(map[int]int)
	for _, v := range arr {
		counts[v]++
	}

	mapCounts := make(map[int]struct{})
	for _, count := range counts {
		mapCounts[count] = struct{}{}
	}

	return len(counts) == len(mapCounts)
}
