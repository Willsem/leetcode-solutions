func longestPalindrome(words []string) int {
	counts := make(map[string]int)
	for _, word := range words {
		counts[word]++
	}

	countToUse := 0
	hasPalindromeForMiddle := false
	for word := range counts {
		reversed := string(word[1]) + string(word[0])
		if word == reversed {
			countToUse += counts[word] - counts[word]%2
			if counts[word]%2 == 1 {
				hasPalindromeForMiddle = true
			}
			continue
		}

		countToUse += min(counts[word], counts[reversed])
	}

	if hasPalindromeForMiddle {
		countToUse++
	}

	return countToUse * 2
}
