func equalSubstring(s string, t string, maxCost int) int {
	left := 0
	currCost, maxLen := 0, 0
	for right := range s {
		currCost += abs(int(s[right]) - int(t[right]))

		for currCost > maxCost {
			currCost -= abs(int(s[left]) - int(t[left]))
			left++
		}

		maxLen = max(maxLen, right-left+1)
	}

	return maxLen
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
