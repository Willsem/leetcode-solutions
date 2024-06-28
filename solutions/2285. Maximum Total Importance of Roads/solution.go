import "sort"

func maximumImportance(n int, roads [][]int) int64 {
	counts := make([]int, n)
	for _, road := range roads {
		counts[road[0]]++
		counts[road[1]]++
	}

	sort.Ints(counts)

	result := int64(0)
	for i, count := range counts {
		result += int64((i + 1) * count)
	}
	return result
}
