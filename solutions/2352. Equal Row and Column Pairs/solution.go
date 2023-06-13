func equalPairs(grid [][]int) int {
	count := 0
	for i := range grid {
		for j := range grid[i] {
			equal := true
			for k := range grid {
				if grid[i][k] != grid[k][j] {
					equal = false
					break
				}
			}

			if equal {
				count++
			}
		}
	}

	return count
}
