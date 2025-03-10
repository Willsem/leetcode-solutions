var vowels = map[byte]struct{}{
	'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {},
}

func isVowel(a byte) bool {
	_, ok := vowels[a]
	return ok
}

func atLeastK(word string, k int) int64 {
	vowelCount := make(map[byte]int)
	consonantCount := 0
	result := 0
	left := 0

	for right := 0; right < len(word); right++ {
		if isVowel(word[right]) {
			vowelCount[word[right]]++
		} else {
			consonantCount++
		}

		for len(vowelCount) == 5 && consonantCount >= k {
			result += len(word) - right

			if isVowel(word[left]) {
				vowelCount[word[left]]--
				if vowelCount[word[left]] <= 0 {
					delete(vowelCount, word[left])
				}
			} else {
				consonantCount--
			}

			left++
		}
	}

	return int64(result)
}

func countOfSubstrings(word string, k int) int64 {
	return atLeastK(word, k) - atLeastK(word, k+1)
}
