func key(l rune) int {
	return int(byte(l) - byte('a'))
}

func makeEqual(words []string) bool {
	counts := make([]int, 26)

	for _, word := range words {
		for _, letter := range word {
			counts[key(letter)]++
		}
	}

	for _, count := range counts {
		if count%len(words) != 0 {
			return false
		}
	}

	return true
}
