func searchMatrix(matrix [][]int, target int) bool {
	row := binarySearch(0, len(matrix), func(i int) bool {
		return matrix[i][0] <= target
	})

	col := binarySearch(0, len(matrix[row]), func(i int) bool {
		return matrix[row][i] <= target
	})

	return matrix[row][col] == target
}

func binarySearch(l, r int, comp func(i int) bool) int {
	for l < r-1 {
		mid := (l + r) / 2
		if comp(mid) {
			l = mid
		} else {
			r = mid
		}
	}

	return l
}
