type coord struct {
	i, j int
}

func firstCompleteIndex(arr []int, mat [][]int) int {
	coords := make(map[int]coord)
	for i := range mat {
		for j, v := range mat[i] {
			coords[v] = coord{i, j}
		}
	}

	countsRow := make(map[int]int)
	countsCol := make(map[int]int)

	for i, v := range arr {
		countsRow[coords[v].i]++
		if countsRow[coords[v].i] == len(mat[0]) {
			return i
		}

		countsCol[coords[v].j]++
		if countsCol[coords[v].j] == len(mat) {
			return i
		}
	}

	return -1
}
