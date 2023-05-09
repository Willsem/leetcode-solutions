const (
	right = iota
	bottom
	left
	top
)

func spiralOrder(matrix [][]int) []int {
	n := len(matrix)
	m := len(matrix[0])

	result := make([]int, n*m)
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}

	i, j, resI := 0, 0, 0
	direction := right
	for resI < len(result) {
		result[resI] = matrix[i][j]
		resI++
		visited[i][j] = true

		switch direction {
		case right:
			j++
			if j >= m || visited[i][j] {
				j--
				i++
				direction = bottom
			}
		case bottom:
			i++
			if i >= n || visited[i][j] {
				i--
				j--
				direction = left
			}
		case left:
			j--
			if j < 0 || visited[i][j] {
				j++
				i--
				direction = top
			}
		case top:
			i--
			if i < 0 || visited[i][j] {
				i++
				j++
				direction = right
			}
		}
	}

	return result
}
