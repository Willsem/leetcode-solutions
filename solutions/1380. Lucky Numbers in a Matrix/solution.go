func luckyNumbers(matrix [][]int) []int {
	mins := make([]int, len(matrix))
	for i := range matrix {
		mins[i] = matrix[i][0]
	}

	maxs := make([]int, len(matrix[0]))
	for j := range matrix[0] {
		maxs[j] = matrix[0][j]
	}

	for i := range matrix {
		for j := range matrix[i] {
			mins[i] = min(mins[i], matrix[i][j])
			maxs[j] = max(maxs[j], matrix[i][j])
		}
	}

	res := make([]int, 0)
	for i := range matrix {
		for j := range matrix[i] {
			if mins[i] == maxs[j] {
				res = append(res, matrix[i][j])
			}
		}
	}
	return res
}
