func countNegatives(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])

	i, j := 0, m-1
	count := 0
	for i < n {
		if grid[i][j] >= 0 {
			count += m - j - 1
			i++
		} else {
			j--
			if j == -1 {
				count += m
				j = 0
				i++
			}
		}
	}

	return count
}
