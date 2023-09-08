func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	res[0] = []int{1}

	for i := 1; i < numRows; i++ {
		res[i] = make([]int, len(res[i-1])+1)
		for j := range res[i] {
			if j == 0 || j == len(res[i])-1 {
				res[i][j] = 1
			} else {
				res[i][j] = res[i-1][j-1] + res[i-1][j]
			}
		}
	}

	return res
}
