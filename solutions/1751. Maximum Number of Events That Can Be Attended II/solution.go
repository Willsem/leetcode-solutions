import "sort"

type Events struct {
	events [][]int
}

func NewEvents(events [][]int) *Events {
	return &Events{
		events: events,
	}
}

func (e *Events) Sort() {
	sort.Slice(e.events, func(i, j int) bool {
		return e.events[i][0] < e.events[j][0]
	})
}

func (e *Events) Count() int {
	return len(e.events)
}

func (e *Events) StartDay(i int) int {
	return e.events[i][0]
}

func (e *Events) EndDay(i int) int {
	return e.events[i][1]
}

func (e *Events) Value(i int) int {
	return e.events[i][2]
}

func maxValue(e [][]int, k int) int {
	events := NewEvents(e)
	events.Sort()

	n := events.Count()
	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for currIndex := n - 1; currIndex >= 0; currIndex-- {
		nextIndex := binarySearch(0, n, func(i int) bool {
			return events.StartDay(i) <= events.EndDay(currIndex)
		})
		for count := 1; count <= k; count++ {
			dp[count][currIndex] =
				max(dp[count][currIndex+1], events.Value(currIndex)+dp[count-1][nextIndex])
		}
	}

	return dp[k][0]
}

func binarySearch(l, r int, comp func(i int) bool) int {
	for l < r {
		mid := (l + r) / 2
		if comp(mid) {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
