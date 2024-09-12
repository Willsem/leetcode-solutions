func countConsistentStrings(allowed string, words []string) int {
	allowedChars := make(map[rune]struct{})
	for _, v := range allowed {
		allowedChars[v] = struct{}{}
	}

	count := 0
	for _, word := range words {
		isCounted := true
		for _, v := range word {
			if _, ok := allowedChars[v]; !ok {
				isCounted = false
				break
			}
		}

		if isCounted {
			count++
		}
	}

	return count
}
