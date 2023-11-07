import (
	"math"
	"sort"
)

func eliminateMaximum(dist []int, speed []int) int {
	arrivalTime := make([]int, len(dist))
	for i := range dist {
		arrivalTime[i] = int(math.Ceil(float64(dist[i]) / float64(speed[i])))
	}

	sort.Ints(arrivalTime)

	for i := range arrivalTime {
		if arrivalTime[i] <= i {
			return i
		}
	}

	return len(arrivalTime)
}
