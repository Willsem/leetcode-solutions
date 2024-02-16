import "sort"

type numCount struct {
	num   int
	count int
}

func findLeastNumOfUniqueInts(arr []int, k int) int {
	counts := make(map[int]int)
	for _, num := range arr {
		counts[num]++
	}

	sortedCounts := make([]numCount, 0, len(counts))
	for num, count := range counts {
		sortedCounts = append(sortedCounts, numCount{num, count})
	}

	sort.Slice(sortedCounts, func(i, j int) bool {
		return sortedCounts[i].count < sortedCounts[j].count
	})

	curr := 0
	for curr < len(sortedCounts) {
		if k >= sortedCounts[curr].count {
			k -= sortedCounts[curr].count
			curr++
		} else {
			break
		}
	}

	return len(sortedCounts) - curr
}
