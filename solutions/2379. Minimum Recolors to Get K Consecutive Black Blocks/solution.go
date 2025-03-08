func minimumRecolors(blocks string, k int) int {
	left := 0
	minWCount, wCount := k, 0
	for right := 0; right < len(blocks); right++ {
		if blocks[right] == 'W' {
			wCount++
		}

		if right-left == k-1 {
			minWCount = min(minWCount, wCount)

			if blocks[left] == 'W' {
				wCount--
			}
			left++
		}
	}

	return minWCount
}
