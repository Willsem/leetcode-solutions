func isVowel(letter byte) bool {
	switch letter {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	}

	return false
}

func maxVowels(s string, k int) int {
	left := 0
	right := k - 1
	vowelCount := 0

	for i := left; i <= right; i++ {
		if isVowel(s[i]) {
			vowelCount++
		}
	}

	maxCount := vowelCount

	for right < len(s)-1 {
		if isVowel(s[left]) {
			vowelCount--
		}

		left++
		right++

		if isVowel(s[right]) {
			vowelCount++
		}

		if vowelCount > maxCount {
			maxCount = vowelCount
		}
	}

	return maxCount
}
