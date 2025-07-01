func possibleStringCount(word string) int {
	count := 1
	var prev rune
	for _, c := range word {
		if c == prev {
			count++
		}
		prev = c
	}

	return count
}
