func islandPerimeter(grid [][]int) int {
	count := 0

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				count += 4

				if i > 0 {
					count -= grid[i-1][j]
				}
				if j > 0 {
					count -= grid[i][j-1]
				}
				if i < len(grid)-1 {
					count -= grid[i+1][j]
				}
				if j < len(grid[i])-1 {
					count -= grid[i][j+1]
				}
			}
		}
	}

	return count
}
