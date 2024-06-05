func commonChars(words []string) []string {
	counts := make(map[rune]int)
	for _, v := range words[0] {
		counts[v]++
	}

	for i := 1; i < len(words); i++ {
		currCounts := make(map[rune]int)
		for _, v := range words[i] {
			currCounts[v]++
		}

		for c := range counts {
			if _, ok := currCounts[c]; !ok {
				delete(counts, c)
			} else {
				counts[c] = min(counts[c], currCounts[c])
			}
		}
	}

	result := make([]string, 0)
	for c, count := range counts {
		for i := 0; i < count; i++ {
			result = append(result, string(c))
		}
	}
	return result
}
