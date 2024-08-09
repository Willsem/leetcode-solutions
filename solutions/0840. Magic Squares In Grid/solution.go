func numMagicSquaresInside(grid [][]int) int {
	if len(grid) < 3 && len(grid[0]) < 3 {
		return 0
	}

	res := 0
	for iLeft, iRight := 0, 2; iRight < len(grid); iLeft, iRight = iLeft+1, iRight+1 {
		for jLeft, jRight := 0, 2; jRight < len(grid[0]); jLeft, jRight = jLeft+1, jRight+1 {
			was := make(map[int]struct{})

			sumRows, sumCols, sumDiags := make([]int, 3), make([]int, 3), make([]int, 2)
			for i := iLeft; i <= iRight; i++ {
				for j := jLeft; j <= jRight; j++ {
					localI, localJ := i-iLeft, j-jLeft

					sumRows[localI] += grid[i][j]
					sumCols[localJ] += grid[i][j]

					if localI == localJ {
						sumDiags[0] += grid[i][j]
					}

					if localI+localJ == 2 {
						sumDiags[1] += grid[i][j]
					}

					was[grid[i][j]] = struct{}{}
				}
			}

			if sumRows[0] != sumRows[1] || sumRows[1] != sumRows[2] {
				continue
			}

			if sumCols[0] != sumCols[1] || sumCols[1] != sumCols[2] {
				continue
			}

			if sumDiags[0] != sumDiags[1] {
				continue
			}

			wasAll := true
			for k := 1; k <= 9; k++ {
				if _, ok := was[k]; !ok {
					wasAll = false
					break
				}
			}

			if wasAll {
				res++
			}
		}
	}

	return res
}
