func setZeroes(matrix [][]int) {
	cols, rows := make(map[int]struct{}), make(map[int]struct{})

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				rows[i] = struct{}{}
				cols[j] = struct{}{}
			}
		}
	}

	for i := range matrix {
		for j := range matrix[i] {
			if _, ok := rows[i]; ok {
				matrix[i][j] = 0
			}
			if _, ok := cols[j]; ok {
				matrix[i][j] = 0
			}
		}
	}
}
