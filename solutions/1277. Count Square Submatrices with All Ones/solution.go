func countSquares(matrix [][]int) int {
	ans := 0

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				continue
			}

			if i == 0 || j == 0 {
				ans += matrix[i][j]
				continue
			}

			matrix[i][j] = min(matrix[i-1][j], matrix[i][j-1], matrix[i-1][j-1]) + 1
			ans += matrix[i][j]
		}
	}

	return ans
}
