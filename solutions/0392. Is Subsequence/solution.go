func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	sPointer := 0
	for i := range t {
		if s[sPointer] == t[i] {
			sPointer++

			if sPointer == len(s) {
				return true
			}
		}
	}

	return sPointer == len(s)
}
