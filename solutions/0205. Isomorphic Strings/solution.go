func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapping := make(map[byte]byte)
	backMapping := make(map[byte]byte)
	for i := range s {
		if m, ok := mapping[s[i]]; ok {
			if m != t[i] {
				return false
			}
		} else {
			mapping[s[i]] = t[i]
		}

		if m, ok := backMapping[t[i]]; ok {
			if m != s[i] {
				return false
			}
		} else {
			backMapping[t[i]] = s[i]
		}
	}

	return true
}
