func queryResults(limit int, queries [][]int) []int {
	result := make([]int, 0, len(queries))
	countColors := make(map[int]int)
	ballToColor := make(map[int]int)

	for _, query := range queries {
		ball, color := query[0], query[1]

		if oldColor, ok := ballToColor[ball]; ok {
			countColors[oldColor]--
			if countColors[oldColor] == 0 {
				delete(countColors, oldColor)
			}
		}

		ballToColor[ball] = color
		countColors[color]++

		result = append(result, len(countColors))
	}

	return result
}
