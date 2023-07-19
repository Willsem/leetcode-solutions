import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	count := 0
	currEnd := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= currEnd {
			currEnd = intervals[i][1]
		} else {
			count++
		}
	}

	return count
}
