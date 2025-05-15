func getLongestSubsequence(words []string, groups []int) []string {
	result := make([]string, 0)
	lastGroup := -1

	for i := range words {
		if groups[i] != lastGroup {
			result = append(result, words[i])
			lastGroup = groups[i]
		}
	}

	return result
}
