func champagneTower(poured int, query_row int, query_glass int) float64 {
	tower := [101][101]float64{}
	tower[0][0] = float64(poured)

	for i := 0; i <= query_row; i++ {
		for j := 0; j <= i; j++ {
			if tower[i-j][j] > 1 {
				tower[i-j+1][j] += (tower[i-j][j] - 1) / 2.0
				tower[i-j][j+1] += (tower[i-j][j] - 1) / 2.0
				tower[i-j][j] = 1.0
			}
		}
	}

	return tower[query_row-query_glass][query_glass]
}
