var directions = [][]int{
	{-2, +1},
	{-2, -1},
	{+2, +1},
	{+2, -1},
	{+1, -2},
	{-1, -2},
	{+1, +2},
	{-1, +2},
}

func knightProbability(n int, k int, row int, column int) float64 {
	prevDp := make([][]float64, n)
	currDp := make([][]float64, n)
	for i := range prevDp {
		prevDp[i] = make([]float64, n)
		currDp[i] = make([]float64, n)
	}

	prevDp[row][column] = 1

	for moves := 1; moves <= k; moves++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				currDp[i][j] = 0

				for _, direction := range directions {
					prevI, prevJ := i+direction[0], j+direction[1]
					if prevI >= 0 && prevI < n && prevJ >= 0 && prevJ < n {
						currDp[i][j] += prevDp[prevI][prevJ] / 8.0
					}
				}
			}
		}

		prevDp, currDp = currDp, prevDp
	}

	probability := 0.0
	for i := range prevDp {
		for j := range prevDp[i] {
			probability += prevDp[i][j]
		}
	}
	return probability
}
