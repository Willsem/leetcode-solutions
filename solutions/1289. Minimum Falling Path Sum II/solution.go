func minFallingPathSum(grid [][]int) int {
	prevMinJ, currMinJ := -1, -1
	for i := range grid {
		if i == 0 {
			for j := range grid[i] {
				if prevMinJ == -1 || grid[i][j] < grid[i][prevMinJ] {
					prevMinJ = j
				}
			}

			continue
		}

		currMinJ = -1
		for j := range grid[i] {
			if j != prevMinJ {
				grid[i][j] += grid[i-1][prevMinJ]
			} else {
				secondMinJ := -1
				for k := range grid[i-1] {
					if k != j && (secondMinJ == -1 || grid[i-1][k] < grid[i-1][secondMinJ]) {
						secondMinJ = k
					}
				}
				grid[i][j] += grid[i-1][secondMinJ]
			}

			if currMinJ == -1 || grid[i][j] < grid[i][currMinJ] {
				currMinJ = j
			}
		}
		prevMinJ = currMinJ
	}

	return grid[len(grid)-1][prevMinJ]
}
