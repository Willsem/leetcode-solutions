import "math"

func minFallingPathSum(matrix [][]int) int {
	ans := math.MaxInt

	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] += minVariant(matrix, i, j)

			if i == len(matrix)-1 && ans > matrix[i][j] {
				ans = matrix[i][j]
			}
		}
	}

	return ans
}

func minVariant(matrix [][]int, i, j int) int {
	if i == 0 {
		return 0
	}

	if j == 0 {
		return min(matrix[i-1][j], matrix[i-1][j+1])
	}

	if j == len(matrix[i])-1 {
		return min(matrix[i-1][j], matrix[i-1][j-1])
	}

	return min(matrix[i-1][j], min(matrix[i-1][j-1], matrix[i-1][j+1]))
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
