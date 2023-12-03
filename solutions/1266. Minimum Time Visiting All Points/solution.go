import "math"

func minTimeToVisitAllPoints(points [][]int) int {
	time := 0
	var prev []int
	for _, point := range points {
		if prev == nil {
			prev = point
			continue
		}

		x := int(math.Abs(float64(point[0] - prev[0])))
		y := int(math.Abs(float64(point[1] - prev[1])))
		time += int(math.Max(float64(x), float64(y)))

		prev = point
	}

	return time
}
