const (
	right = iota
	bottom
	left
	top
)

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	i, j, resI := 0, 0, 1
	direction := right
	for resI <= n*n {
		matrix[i][j] = resI
		resI++

		switch direction {
		case right:
			j++
			if j >= n || matrix[i][j] > 0 {
				j--
				i++
				direction = bottom
			}
		case bottom:
			i++
			if i >= n || matrix[i][j] > 0 {
				i--
				j--
				direction = left
			}
		case left:
			j--
			if j < 0 || matrix[i][j] > 0 {
				j++
				i--
				direction = top
			}
		case top:
			i--
			if i < 0 || matrix[i][j] > 0 {
				i++
				j++
				direction = right
			}
		}
	}

	return matrix
}
