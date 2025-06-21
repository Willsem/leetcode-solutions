func minimumDeletions(word string, k int) int {
	freq := make(map[byte]int)
	for i := range word {
		freq[word[i]]++
	}

	minDeleted := len(word)
	for _, baseCount := range freq {
		deleted := 0
		for _, f := range freq {
			if baseCount > f {
				deleted += f
			} else if f-baseCount > k {
				deleted += f - baseCount - k
			}
		}
		minDeleted = min(minDeleted, deleted)
	}

	return minDeleted
}
