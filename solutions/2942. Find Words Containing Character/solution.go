func findWordsContaining(words []string, x byte) []int {
	result := make([]int, 0)
	for i, word := range words {
		for _, c := range word {
			if byte(c) == x {
				result = append(result, i)
				break
			}
		}
	}
	return result
}
