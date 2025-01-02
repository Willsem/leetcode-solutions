func isVowel(a byte) bool {
	switch a {
	case 'e', 'u', 'i', 'o', 'a':
		return true
	}

	return false
}

func vowelStrings(words []string, queries [][]int) []int {
	prefixCount := make([]int, len(words))

	count := 0
	for i, word := range words {
		if isVowel(word[0]) && isVowel(word[len(word)-1]) {
			count++
		}
		prefixCount[i] = count
	}

	result := make([]int, len(queries))
	for i, query := range queries {
		var c0 int
		if query[0] > 0 {
			c0 = prefixCount[query[0]-1]
		}
		c1 := prefixCount[query[1]]

		result[i] = c1 - c0
	}
	return result
}
