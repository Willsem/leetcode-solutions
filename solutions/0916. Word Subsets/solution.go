func wordSubsets(words1 []string, words2 []string) []string {
	maxCounts2 := make(map[rune]int)
	for _, word := range words2 {
		counts := calcCounts(word)
		for v, count := range counts {
			maxCounts2[v] = max(maxCounts2[v], count)
		}
	}

	result := make([]string, 0)
	for _, word := range words1 {
		counts := calcCounts(word)
		isUniverse := true
		for v, count := range maxCounts2 {
			if count > counts[v] {
				isUniverse = false
				break
			}
		}

		if isUniverse {
			result = append(result, word)
		}
	}

	return result
}

func calcCounts(word string) map[rune]int {
	counts := make(map[rune]int)
	for _, v := range word {
		counts[v]++
	}
	return counts
}
