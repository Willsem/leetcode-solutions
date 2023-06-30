func latestDayToCross(row int, col int, cells [][]int) int {
	grid := make([][]bool, row)
	for i := 0; i < row; i++ {
		grid[i] = make([]bool, col)
	}

	left, right := 0, len(cells)-1
	for left < right {
		mid := (left + right) / 2
		for i := left; i <= mid; i++ {
			grid[cells[i][0]-1][cells[i][1]-1] = true
		}
		if !canReachEnd(row, col, grid) {
			for i := left; i <= mid; i++ {
				grid[cells[i][0]-1][cells[i][1]-1] = false
			}
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func canReachEnd(row int, col int, grid [][]bool) bool {
	visited := make([][]bool, row)
	for i := 0; i < row; i++ {
		visited[i] = make([]bool, col)
	}

	var dfs func(y, x int) bool
	dfs = func(y, x int) bool {
		visited[y][x] = true
		if y == row-1 {
			return true
		}
		if y < row-1 && !grid[y+1][x] && !visited[y+1][x] && dfs(y+1, x) {
			return true
		}
		if x < col-1 && !grid[y][x+1] && !visited[y][x+1] && dfs(y, x+1) {
			return true
		}
		if x > 0 && !grid[y][x-1] && !visited[y][x-1] && dfs(y, x-1) {
			return true
		}
		if y > 0 && !grid[y-1][x] && !visited[y-1][x] && dfs(y-1, x) {
			return true
		}
		return false
	}
	for i := 0; i < len(grid[0]); i++ {
		if grid[0][i] {
			visited[0][i] = true
			continue
		}
		if dfs(0, i) {
			return true
		}
	}
	return false
}
