func findMissingAndRepeatedValues(grid [][]int) []int {
	values := make(map[int]struct{})
	res := []int{0, 0}
	sum := 0
	for _, row := range grid {
		for _, v := range row {
			if _, ok := values[v]; ok {
				res[0] = v
			}
			values[v] = struct{}{}
			sum += v
		}
	}

	sum -= res[0]

	n := len(grid) * len(grid)
	res[1] = n*(1+n)/2 - sum

	return res
}
