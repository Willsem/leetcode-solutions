func maximumGain(s string, x int, y int) int {
	if x < y {
		x, y = y, x
		s = reverse(s)
	}

	aCount, bCount, points := 0, 0, 0
	for _, char := range s {
		if char == 'a' {
			aCount++
		} else if char == 'b' {
			if aCount > 0 {
				aCount--
				points += x
			} else {
				bCount++
			}
		} else {
			points += min(bCount, aCount) * y
			aCount, bCount = 0, 0
		}
	}

	points += min(bCount, aCount) * y
	return points
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
