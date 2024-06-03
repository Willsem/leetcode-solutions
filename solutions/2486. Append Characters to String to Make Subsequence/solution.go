func appendCharacters(s string, t string) int {
	ti := 0
	for si := 0; si < len(s); si++ {
		if s[si] == t[ti] {
			ti++
			if ti == len(t) {
				return 0
			}
		}
	}

	return len(t) - ti
}
