func diagonalSum(mat [][]int) int {
	sum := 0

	if len(mat)%2 != 0 {
		centerIndex := (len(mat) - 1) / 2
		sum = -mat[centerIndex][centerIndex]
	}

	for i := range mat {
		sum += mat[i][i] + mat[i][len(mat)-1-i]
	}

	return sum
}
