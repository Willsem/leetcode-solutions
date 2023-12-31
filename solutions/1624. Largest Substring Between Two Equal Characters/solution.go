func maxLengthBetweenEqualCharacters(s string) int {
	firstChar := make(map[rune]int)
	max := 0
	updated := false
	for i, c := range s {
		if _, ok := firstChar[c]; !ok {
			firstChar[c] = i
		}

		prev := firstChar[c]
		diff := i - prev - 1
		if diff >= max {
			max = diff
			updated = true
		}
	}

	if !updated {
		return -1
	}

	return max
}
