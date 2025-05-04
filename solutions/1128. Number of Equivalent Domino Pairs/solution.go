func numEquivDominoPairs(dominoes [][]int) int {
	counts := make(map[int]int)
	count := 0

	for _, d := range dominoes {
		if d[0] > d[1] {
			d[0], d[1] = d[1], d[0]
		}

		value := d[0]*10 + d[1]
		count += counts[value]
		counts[value]++
	}

	return count
}
