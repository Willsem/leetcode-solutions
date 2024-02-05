func firstUniqChar(s string) int {
	counts := make(map[rune]int)
	for _, c := range s {
		counts[c]++
	}

	for i, c := range s {
		if counts[c] == 1 {
			return i
		}
	}

	return -1
}
