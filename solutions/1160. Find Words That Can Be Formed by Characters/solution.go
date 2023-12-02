func countCharacters(words []string, chars string) int {
	counts := make(map[rune]int)
	for _, char := range chars {
		counts[char]++
	}

	sum := 0
	for _, word := range words {
		counts := copy(counts)
		wasBreak := false
		for _, c := range word {
			if _, ok := counts[c]; !ok {
				wasBreak = true
				break
			}

			counts[c]--
			if counts[c] == 0 {
				delete(counts, c)
			}
		}

		if !wasBreak {
			sum += len(word)
		}
	}

	return sum
}

func copy(m map[rune]int) map[rune]int {
	new := make(map[rune]int, len(m))
	for k, v := range m {
		new[k] = v
	}
	return new
}
