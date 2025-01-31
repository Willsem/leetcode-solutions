var dirs = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func largestIsland(grid [][]int) int {
	n := len(grid)
	islandID := 2
	islandSize := make(map[int]int)

	var dfs func(x, y, id int) int
	dfs = func(x, y, id int) int {
		if !isValid(x, y, n) || grid[x][y] != 1 {
			return 0
		}

		grid[x][y] = id
		size := 1
		for _, dir := range dirs {
			size += dfs(x+dir[0], y+dir[1], id)
		}
		return size
	}

	maxIsland := 0
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			if grid[x][y] == 1 {
				islandSize[islandID] = dfs(x, y, islandID)
				maxIsland = max(maxIsland, islandSize[islandID])
				islandID++
			}
		}
	}

	hasZero := false
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			if grid[x][y] == 0 {
				hasZero = true

				seen := make(map[int]struct{})
				newSize := 1
				for _, dir := range dirs {
					nx, ny := x+dir[0], y+dir[1]
					if isValid(nx, ny, n) {
						id := grid[nx][ny]

						size, ok := islandSize[id]
						_, was := seen[id]
						if ok && !was {
							seen[id] = struct{}{}
							newSize += size
						}
					}
				}

				maxIsland = max(maxIsland, newSize)
			}
		}
	}

	if !hasZero {
		return n * n
	}
	return maxIsland
}

func isValid(x, y, n int) bool {
	return x >= 0 && y >= 0 && x < n && y < n
}
