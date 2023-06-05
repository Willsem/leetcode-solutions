func checkStraightLine(coordinates [][]int) bool {
	x1 := coordinates[0][0]
	y1 := coordinates[0][1]

	x2 := coordinates[1][0]
	y2 := coordinates[1][1]

	if x1 == x2 {
		for i := 2; i < len(coordinates); i++ {
			x := coordinates[i][0]

			if x != x1 {
				return false
			}
		}

		return true
	}

	k := float64(y1-y2) / float64(x1-x2)
	b := float64(y1) - k*float64(x1)

	for i := 2; i < len(coordinates); i++ {
		x := float64(coordinates[i][0])
		y := float64(coordinates[i][1])

		if y != k*x+b {
			return false
		}
	}

	return true
}
