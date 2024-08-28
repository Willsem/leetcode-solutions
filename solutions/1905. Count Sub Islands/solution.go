func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	n, m := len(grid1), len(grid1[0])

	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}

	var dfs func(i, j int, isCounted bool) bool
	dfs = func(i, j int, isCounted bool) bool {
		if i < 0 || i >= n || j < 0 || j >= m || visited[i][j] {
			return isCounted
		}
		if grid2[i][j] == 0 {
			return isCounted
		}

		visited[i][j] = true

		if grid1[i][j] == 0 {
			isCounted = false
		}

		isCounted1 := dfs(i+1, j, isCounted)
		isCounted2 := dfs(i-1, j, isCounted)
		isCounted3 := dfs(i, j+1, isCounted)
		isCounted4 := dfs(i, j-1, isCounted)

		return isCounted1 && isCounted2 && isCounted3 && isCounted4
	}

	count := 0
	for i := range grid2 {
		for j := range grid2[i] {
			if grid2[i][j] == 1 && !visited[i][j] {
				if dfs(i, j, true) {
					count++
				}
			}
		}
	}

	return count
}
