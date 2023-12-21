import "sort"

func maxWidthOfVerticalArea(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] > points[j][0]
	})

	prevX := points[0][0]
	max := 0
	for _, point := range points {
		if prevX-point[0] > max {
			max = prevX - point[0]
		}
		prevX = point[0]
	}

	return max
}
