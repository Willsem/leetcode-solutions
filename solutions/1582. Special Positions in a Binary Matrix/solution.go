func numSpecial(mat [][]int) int {
	countI, countJ := make([]int, len(mat)), make([]int, len(mat[0]))

	for i := range mat {
		for j := range mat[i] {
			if mat[i][j] == 1 {
				countI[i]++
				countJ[j]++
			}
		}
	}

	res := 0
	for i := range mat {
		for j := range mat[i] {
			if mat[i][j] == 1 && countI[i] == 1 && countJ[j] == 1 {
				res++
			}
		}
	}

	return res
}
