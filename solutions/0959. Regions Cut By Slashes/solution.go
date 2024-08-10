func regionsBySlashes(grid []string) int {
	matrix := make([][]byte, len(grid)*3)
	for i := range matrix {
		matrix[i] = make([]byte, len(grid[0])*3)
	}

	for i, row := range grid {
		for j, ceil := range row {
			fillCeil(matrix, i, j, ceil)
		}
	}

	count := 0
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				count++
				bfs(matrix, i, j)
			}
		}
	}

	return count
}

func fillCeil(matrix [][]byte, i, j int, symbol rune) {
	if symbol == ' ' {
		return
	}

	for fillI := i * 3; fillI < i*3+3; fillI++ {
		for fillJ := j * 3; fillJ < j*3+3; fillJ++ {
			localI, localJ := fillI-i*3, fillJ-j*3
			if symbol == '/' && localI+localJ == 2 || symbol == '\\' && localI == localJ {
				matrix[fillI][fillJ] = 1
			}
		}
	}
}

func bfs(matrix [][]byte, i, j int) {
	if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[0]) {
		return
	}

	if matrix[i][j] != 0 {
		return
	}

	matrix[i][j] = 2

	bfs(matrix, i+1, j)
	bfs(matrix, i-1, j)
	bfs(matrix, i, j+1)
	bfs(matrix, i, j-1)
}
