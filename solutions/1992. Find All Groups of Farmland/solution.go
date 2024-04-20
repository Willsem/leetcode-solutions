func findFarmland(land [][]int) [][]int {
	res := make([][]int, 0)

	for i := range land {
		for j := range land[i] {
			if land[i][j] == 1 {
				x, y := 0, 0
				for x = i; x < len(land) && land[x][j] == 1; x++ {
					for y = j; y < len(land[x]) && land[x][y] == 1; y++ {
						land[x][y] = 0
					}
				}

				res = append(res, []int{i, j, x - 1, y - 1})
			}
		}
	}

	return res
}
