func numberOfAlternatingGroups(colors []int, k int) int {
	result := 0
	altElCount := 1
	lastColor := colors[0]

	for i := 1; i < len(colors)+k-1; i++ {
		index := i % len(colors)

		if colors[index] == lastColor {
			altElCount = 1
			lastColor = colors[index]
			continue
		}

		altElCount++

		if altElCount >= k {
			result++
		}

		lastColor = colors[index]
	}

	return result
}
