func restoreMatrix(rowSum []int, colSum []int) [][]int {
	n, m := len(rowSum), len(colSum)
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, m)
	}

	i, j := 0, 0
	for i < n && j < m {
		matrix[i][j] = min(rowSum[i], colSum[j])
		rowSum[i] -= matrix[i][j]
		colSum[j] -= matrix[i][j]

		if rowSum[i] == 0 {
			i++
		} else {
			j++
		}
	}

	return matrix
}
