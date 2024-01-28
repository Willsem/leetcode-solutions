func numSubmatrixSumTarget(matrix [][]int, target int) int {
	n, m := len(matrix), len(matrix[0])

	for i := 0; i < n; i++ {
		for j := 1; j < m; j++ {
			matrix[i][j] += matrix[i][j-1]
		}
	}

	count := 0
	for i := 0; i < m; i++ {
		for j := i; j < m; j++ {
			m := make(map[int]int)
			m[0] = 1
			sum := 0

			for row := 0; row < n; row++ {
				temp := 0
				if i > 0 {
					temp = matrix[row][i-1]
				}
				sum += matrix[row][j] - temp
				count += m[sum-target]
				m[sum]++
			}
		}
	}

	return count
}
