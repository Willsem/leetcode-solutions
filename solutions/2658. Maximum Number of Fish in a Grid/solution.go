func findMaxFish(grid [][]int) int {
	currFish := 0

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || grid[i][j] == 0 {
			return
		}

		currFish += grid[i][j]
		grid[i][j] = 0

		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}

	maxFish := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] > 0 {
				currFish = 0
				dfs(i, j)
				maxFish = max(maxFish, currFish)
			}
		}
	}

	return maxFish
}
