func numberOfSubstrings(s string) int {
	counts := make(map[byte]int)
	result := 0

	left := 0
	for right := 0; right < len(s); right++ {
		counts[s[right]]++

		for len(counts) == 3 {
			result += len(s) - right
			counts[s[left]]--
			if counts[s[left]] == 0 {
				delete(counts, s[left])
			}
			left++
		}
	}

	return result
}
