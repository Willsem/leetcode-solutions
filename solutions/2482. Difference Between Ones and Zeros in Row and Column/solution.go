func onesMinusZeros(grid [][]int) [][]int {
	n := len(grid)
	m := len(grid[0])
	col := make([]int, m)
	row := make([]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				row[i]++
				col[j]++
			} else {
				row[i]--
				col[j]--
			}
		}
	}

	jaggedArray := make([][]int, n)

	for i := 0; i < n; i++ {
		jaggedArray[i] = make([]int, m)
		for j := 0; j < m; j++ {
			jaggedArray[i][j] = col[j] + row[i]
		}
	}

	return jaggedArray
}
