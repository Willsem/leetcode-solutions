func largestLocal(grid [][]int) [][]int {
	result := make([][]int, len(grid)-2)
	for i := range result {
		result[i] = make([]int, len(grid[i])-2)
		for j := range result[i] {
			result[i][j] = findMax(grid, i, j)
		}
	}
	return result
}

func findMax(grid [][]int, i, j int) int {
	m := grid[i][j]
	for _, gridI := range []int{i, i + 1, i + 2} {
		for _, gridJ := range []int{j, j + 1, j + 2} {
			m = max(m, grid[gridI][gridJ])
		}
	}
	return m
}
