func calculateWordScore(word string, score []int) int {
	total := 0
	for _, char := range word {
		total += score[char-'a']
	}
	return total
}

func generateSubsets(words []string) [][]string {
	n := len(words)
	allSubsets := [][]string{}
	for i := 0; i < (1 << n); i++ {
		subset := []string{}
		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 {
				subset = append(subset, words[j])
			}
		}
		allSubsets = append(allSubsets, subset)
	}
	return allSubsets
}

func maxScoreWords(words []string, letters []byte, score []int) int {
	letterCount := make(map[byte]int)
	for _, letter := range letters {
		letterCount[letter]++
	}

	maxScore := 0
	allSubsets := generateSubsets(words)

	for _, subset := range allSubsets {
		subsetCount := make(map[byte]int)
		subsetScore := 0
		valid := true

		for _, word := range subset {
			wordCount := make(map[byte]int)
			for _, char := range []byte(word) {
				wordCount[char]++
				subsetCount[char]++
				if subsetCount[char] > letterCount[char] {
					valid = false
					break
				}
			}
			if !valid {
				break
			}
			subsetScore += calculateWordScore(word, score)
		}

		if valid {
			if subsetScore > maxScore {
				maxScore = subsetScore
			}
		}
	}

	return maxScore
}
