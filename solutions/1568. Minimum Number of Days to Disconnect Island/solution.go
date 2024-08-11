func minDays(grid [][]int) int {
	if isDisconnected(grid) {
		return 0
	}

	m, n := len(grid), len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				grid[i][j] = 0
				if isDisconnected(grid) {
					return 1
				}
				grid[i][j] = 1
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				grid[i][j] = 0
				for x := 0; x < m; x++ {
					for y := 0; y < n; y++ {
						if grid[x][y] == 1 {
							grid[x][y] = 0
							if isDisconnected(grid) {
								return 2
							}
							grid[x][y] = 1
						}
					}
				}
				grid[i][j] = 1
			}
		}
	}
	return 2
}

func isDisconnected(grid [][]int) bool {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	landCount := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				landCount++
				if !visited[i][j] {
					if landCount > 1 {
						return true
					}
					bfs(grid, visited, i, j)
				}
			}
		}
	}
	return landCount == 0
}

func bfs(grid [][]int, visited [][]bool, i, j int) {
	m, n := len(grid), len(grid[0])
	dirX := []int{-1, 1, 0, 0}
	dirY := []int{0, 0, -1, 1}

	queue := [][]int{{i, j}}
	visited[i][j] = true

	for len(queue) > 0 {
		x, y := queue[0][0], queue[0][1]
		queue = queue[1:]

		for d := 0; d < 4; d++ {
			newX, newY := x+dirX[d], y+dirY[d]
			if newX >= 0 && newX < m && newY >= 0 && newY < n && grid[newX][newY] == 1 && !visited[newX][newY] {
				visited[newX][newY] = true
				queue = append(queue, []int{newX, newY})
			}
		}
	}
}
