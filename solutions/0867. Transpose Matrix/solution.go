func transpose(matrix [][]int) [][]int {
	res := make([][]int, len(matrix[0]))
	for i := range matrix {
		for j := range matrix[i] {
			if i == 0 {
				res[j] = make([]int, len(matrix))
			}

			res[j][i] = matrix[i][j]
		}
	}
	return res
}
