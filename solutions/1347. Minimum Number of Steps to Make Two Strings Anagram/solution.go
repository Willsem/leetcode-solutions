func minSteps(s string, t string) int {
	count := make(map[byte]int)

	for i := range s {
		count[s[i]]--
		count[t[i]]++
	}

	ans := 0
	for _, c := range count {
		if c > 0 {
			ans += c
		}
	}

	return ans
}
