func longestStrChain(words []string) int {
	wordSet := make(map[string]struct{}, len(words))
	for _, word := range words {
		wordSet[word] = struct{}{}
	}

	memo := make(map[string]int, 0)

	var dfs func(word string) int
	dfs = func(word string) int {
		if _, ok := wordSet[word]; !ok {
			return 0
		}

		if res, ok := memo[word]; ok {
			return res
		}

		maxChain := 1
		for i := range word {
			nextWord := word[:i] + word[i+1:]
			chain := 1 + dfs(nextWord)

			if chain > maxChain {
				maxChain = chain
			}
		}

		memo[word] = maxChain
		return maxChain
	}

	result := dfs(words[0])
	for i := 1; i < len(words); i++ {
		currResult := dfs(words[i])
		if currResult > result {
			result = currResult
		}
	}

	return result
}
