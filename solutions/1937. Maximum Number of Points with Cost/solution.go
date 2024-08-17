func maxPoints(points [][]int) int64 {
	n := len(points)
	m := len(points[0])
	prevRow := points[0]

	for i := 1; i < n; i++ {
		leftMax := make([]int, m)
		rightMax := make([]int, m)
		currRow := make([]int, m)

		leftMax[0] = prevRow[0]
		for j := 1; j < m; j++ {
			leftMax[j] = max(leftMax[j-1]-1, prevRow[j])
		}

		rightMax[m-1] = prevRow[m-1]
		for j := m - 2; j >= 0; j-- {
			rightMax[j] = max(rightMax[j+1]-1, prevRow[j])
		}

		for j := 0; j < m; j++ {
			currRow[j] = points[i][j] + max(leftMax[j], rightMax[j])
		}

		prevRow = currRow
	}

	res := prevRow[0]
	for i := 1; i < len(prevRow); i++ {
		res = max(res, prevRow[i])
	}
	return int64(res)
}
