func isAnagram(s string, t string) bool {
	letters := make(map[rune]int)
	for _, c := range s {
		letters[c]++
	}

	for _, c := range t {
		if _, ok := letters[c]; !ok {
			return false
		}

		letters[c]--
		if letters[c] == 0 {
			delete(letters, c)
		}
	}

	return len(letters) == 0
}
