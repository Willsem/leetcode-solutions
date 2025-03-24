import "sort"

func countDays(days int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		if meetings[i][0] == meetings[j][0] {
			return meetings[i][1] < meetings[j][1]
		}
		return meetings[i][0] < meetings[j][0]
	})

	merged := make([][]int, 0)
	currStart := meetings[0][0]
	prevFinish := meetings[0][1]
	for i := 1; i < len(meetings); i++ {
		if prevFinish >= meetings[i][0]-1 {
			prevFinish = max(prevFinish, meetings[i][1])
			continue
		}

		merged = append(merged, []int{currStart, prevFinish})
		currStart = meetings[i][0]
		prevFinish = meetings[i][1]
	}
	merged = append(merged, []int{currStart, prevFinish})

	count := 0
	prevFinish = 0
	for _, busy := range merged {
		currStart := busy[0]
		count += currStart - prevFinish - 1
		prevFinish = busy[1]
	}
	count += days - prevFinish

	return count
}
