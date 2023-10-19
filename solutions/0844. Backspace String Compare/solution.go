func backspaceCompare(s string, t string) bool {
	sCounter, tCounter := 0, 0
	sPointer, tPointer := len(s)-1, len(t)-1

	for sPointer >= 0 || tPointer >= 0 {
		for sPointer >= 0 && s[sPointer] == '#' {
			sCounter++
			sPointer--
		}
		for sPointer >= 0 && sCounter > 0 && s[sPointer] != '#' {
			sCounter--
			sPointer--
		}

		for tPointer >= 0 && t[tPointer] == '#' {
			tCounter++
			tPointer--
		}
		for tPointer >= 0 && tCounter > 0 && t[tPointer] != '#' {
			tCounter--
			tPointer--
		}

		if tCounter == 0 && sCounter == 0 && sPointer >= 0 && tPointer >= 0 {
			if s[sPointer] == '#' || t[tPointer] == '#' {
				continue
			}

			if s[sPointer] != t[tPointer] {
				return false
			}

			sPointer--
			tPointer--
		} else {
			if sPointer == -1 && tPointer >= 0 && tCounter == 0 && t[tPointer] != '#' {
				return false
			}
			if tPointer == -1 && sPointer >= 0 && sCounter == 0 && s[sPointer] != '#' {
				return false
			}
		}
	}

	return sPointer == -1 && tPointer == -1
}
