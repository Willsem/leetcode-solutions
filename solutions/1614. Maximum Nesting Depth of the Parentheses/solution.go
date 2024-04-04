func maxDepth(s string) int {
	counter, maxCount := 0, 0
	for _, c := range s {
		if c == '(' {
			counter++
		} else if c == ')' {
			counter--
		}

		maxCount = max(maxCount, counter)
	}
	return maxCount
}
