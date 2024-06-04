func longestPalindrome(s string) int {
	counts := make(map[rune]int)
	for _, c := range s {
		counts[c]++
	}

	count := len(s)
	wasOdd := false
	for _, v := range counts {
		if v%2 == 1 {
			if !wasOdd {
				wasOdd = true
			} else {
				count--
			}
		}
	}

	return count
}
