func maxDistance(s string, k int) int {
	x, y := 0, 0
	result := 0

	for i, c := range s {
		switch c {
		case 'N':
			y++
		case 'S':
			y--
		case 'E':
			x++
		case 'W':
			x--
		}

		result = max(result, min(abs(x)+abs(y)+k*2, i+1))
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
