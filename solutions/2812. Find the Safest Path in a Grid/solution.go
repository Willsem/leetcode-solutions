func dfs(x, y int, gridP *[][]int, visitedP *[][]bool) bool {
	grid := *gridP
	visited := *visitedP
	if x == len(grid)-1 && x == y {
		return true
	}
	visited[x][y] = true
	if x+1 < len(grid) && !visited[x+1][y] && grid[x+1][y] != 1 {
		if dfs(x+1, y, gridP, visitedP) {
			return true
		}
	}
	if y+1 < len(grid) && !visited[x][y+1] && grid[x][y+1] != 1 {
		if dfs(x, y+1, gridP, visitedP) {
			return true
		}
	}
	if x-1 >= 0 && !visited[x-1][y] && grid[x-1][y] != 1 {
		if dfs(x-1, y, gridP, visitedP) {
			return true
		}
	}
	if y-1 >= 0 && !visited[x][y-1] && grid[x][y-1] != 1 {
		if dfs(x, y-1, gridP, visitedP) {
			return true
		}
	}
	return false
}

func maximumSafenessFactor(grid [][]int) int {
	n := len(grid)
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	ans := 0
	if grid[0][0] != 1 && dfs(0, 0, &grid, &visited) {
		ans++
	}

	for true {
		changes := make([][2]int, 0)
		for i := range grid {
			for j := range grid[i] {
				if grid[i][j] == 1 {
					continue
				}
				if i-1 >= 0 && grid[i-1][j] == 1 {
					changes = append(changes, [2]int{i, j})
				}
				if j-1 >= 0 && grid[i][j-1] == 1 {
					changes = append(changes, [2]int{i, j})
				}
				if i+1 < n && grid[i+1][j] == 1 {
					changes = append(changes, [2]int{i, j})
				}
				if j+1 < n && grid[i][j+1] == 1 {
					changes = append(changes, [2]int{i, j})
				}
			}
		}
		for _, change := range changes {
			grid[change[0]][change[1]] = 1
		}
		for i := range visited {
			visited[i] = make([]bool, n)
		}
		if grid[0][0] != 1 && dfs(0, 0, &grid, &visited) {
			ans++
		} else {
			break
		}
	}
	return ans
}
