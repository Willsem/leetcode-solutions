func partitionLabels(s string) []int {
	lastAppear := make(map[byte]int)
	for i := range s {
		lastAppear[s[i]] = i
	}

	result := make([]int, 0)
	start, end := 0, 0
	firstAppear := make(map[byte]int)
	for i := range s {
		if _, ok := firstAppear[s[i]]; !ok {
			firstAppear[s[i]] = i
		}

		if end < firstAppear[s[i]] {
			result = append(result, end-start+1)
			start = i
			end = i
		}

		end = max(end, lastAppear[s[i]])
	}

	if end-start+1 > 0 {
		result = append(result, end-start+1)
	}

	return result
}
