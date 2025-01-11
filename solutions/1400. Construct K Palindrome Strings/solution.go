func canConstruct(s string, k int) bool {
	if k > len(s) {
		return false
	}

	if k == len(s) {
		return true
	}

	counts := make(map[rune]int)
	for _, v := range s {
		counts[v]++
	}

	oddCount := 0
	for _, count := range counts {
		if count%2 == 1 {
			oddCount++
		}
	}

	return oddCount <= k
}
