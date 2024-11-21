const (
	UNGUARDED = iota
	GUARDED
	GUARD
	WALL
)

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	fillEntity(matrix, guards, GUARD)
	fillEntity(matrix, walls, WALL)

	for _, guard := range guards {
		x, y := guard[0], guard[1]

		// go down
		for i, j := x+1, y; i < m; i++ {
			if matrix[i][j] == GUARD || matrix[i][j] == WALL {
				break
			}
			matrix[i][j] = GUARDED
		}

		// go up
		for i, j := x-1, y; i >= 0; i-- {
			if matrix[i][j] == GUARD || matrix[i][j] == WALL {
				break
			}
			matrix[i][j] = GUARDED
		}

		// go left
		for i, j := x, y-1; j >= 0; j-- {
			if matrix[i][j] == GUARD || matrix[i][j] == WALL {
				break
			}
			matrix[i][j] = GUARDED
		}

		// go right
		for i, j := x, y+1; j < n; j++ {
			if matrix[i][j] == GUARD || matrix[i][j] == WALL {
				break
			}
			matrix[i][j] = GUARDED
		}
	}

	count := 0
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == UNGUARDED {
				count++
			}
		}
	}

	return count
}

func fillEntity(matrix [][]int, points [][]int, entity int) {
	for _, point := range points {
		matrix[point[0]][point[1]] = entity
	}
}
