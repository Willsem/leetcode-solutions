import "sort"

func maxTwoEvents(events [][]int) int {
	sort.Slice(events, func(i, j int) bool {
		return events[i][1] > events[j][1]
	})

	n := len(events)
	maxVals := make([]int, n+1)
	maxVals[n] = 0

	for i := n - 1; i >= 0; i-- {
		maxVals[i] = max(maxVals[i+1], events[i][2])
	}

	best := 0

	for i := 0; i < n; i++ {
		currentValue := events[i][2]
		idx := sort.Search(n, func(j int) bool {
			return events[j][1] < events[i][0]
		})
		if idx < n {
			currentValue += maxVals[idx]
		}
		best = max(best, currentValue)
	}

	return best
}
