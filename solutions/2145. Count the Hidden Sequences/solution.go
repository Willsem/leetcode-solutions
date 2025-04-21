func numberOfArrays(differences []int, lower int, upper int) int {
	x, y, curr := 0, 0, 0
	for _, diff := range differences {
		curr += diff
		x, y = min(x, curr), max(y, curr)
		if y-x > upper-lower {
			return 0
		}
	}
	return (upper - lower) - (y - x) + 1
}
