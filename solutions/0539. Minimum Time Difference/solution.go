import (
	"strconv"
	"strings"
)

const maxMinutes = 24 * 60

func findMinDifference(timePoints []string) int {
	minutesMap := make(map[int]struct{})
	for _, point := range timePoints {
		t := parseTime(point)
		if _, ok := minutesMap[t]; ok {
			return 0
		}

		minutesMap[t] = struct{}{}
	}

	prev := -1
	minDiff := maxMinutes
	first := -1
	for t := 0; t < maxMinutes; t++ {
		if _, ok := minutesMap[t]; !ok {
			continue
		}

		if first == -1 {
			first = t
		}

		if prev != -1 {
			minDiff = min(minDiff, t-prev)
		}

		prev = t
	}

	return min(minDiff, maxMinutes-prev+first)
}

func parseTime(time string) int {
	split := strings.Split(time, ":")
	hours, _ := strconv.Atoi(split[0])
	minutes, _ := strconv.Atoi(split[1])
	return hours*60 + minutes
}
