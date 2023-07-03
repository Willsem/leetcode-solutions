func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	if s == goal {
		freq := make(map[rune]struct{}, len(s))
		for _, c := range s {
			if _, ok := freq[c]; !ok {
				freq[c] = struct{}{}
			} else {
				return true
			}
		}

		return false
	}

	firstIndex := -1
	secondIndex := -1
	for i := range s {
		if s[i] != goal[i] {
			if firstIndex == -1 {
				firstIndex = i
			} else if secondIndex == -1 {
				secondIndex = i
			} else {
				return false
			}
		}
	}

	if secondIndex == -1 {
		return false
	}

	return s[firstIndex] == goal[secondIndex] && s[secondIndex] == goal[firstIndex]
}
