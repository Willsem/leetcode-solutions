func minimumLength(s string) int {
	counts := make(map[rune]int)
	for _, v := range s {
		counts[v]++
	}

	result := 0
	for _, count := range counts {
		if count%2 == 0 {
			result += 2
		} else {
			result += 1
		}
	}

	return result
}
