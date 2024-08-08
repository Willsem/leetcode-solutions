var directions [][]int = [][]int{
	{0, 1}, {1, 0}, {0, -1}, {-1, 0},
}

func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
	matrix := make([][]int, 0)

	x, y := rStart, cStart
	for step, dir := 1, 0; len(matrix) < rows*cols; step++ {
		for i := 0; i < 2; i++ {
			for j := 0; j < step; j++ {
				if x >= 0 && x < rows && y >= 0 && y < cols {
					matrix = append(matrix, []int{x, y})
				}

				x += directions[dir][0]
				y += directions[dir][1]
			}

			dir = (dir + 1) % 4
		}
	}

	return matrix
}
