func countServers(grid [][]int) int {
	rows := make([]int, len(grid))
	cols := make([]int, len(grid[0]))

	for i := range grid {
		for j := range grid[i] {
			rows[i] += grid[i][j]
			cols[j] += grid[i][j]
		}
	}

	count := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 && (rows[i] > 1 || cols[j] > 1) {
				count++
			}
		}
	}

	return count
}
