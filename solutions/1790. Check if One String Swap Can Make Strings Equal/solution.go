func areAlmostEqual(s1 string, s2 string) bool {
	wasNotEq := -1
	wasSecondNotEq := false
	for i := range s1 {
		if s1[i] != s2[i] {
			if wasNotEq != -1 {
				if wasSecondNotEq || s1[wasNotEq] != s2[i] || s2[wasNotEq] != s1[i] {
					return false
				}

				wasSecondNotEq = true
			} else {
				wasNotEq = i
			}
		}
	}

	return wasNotEq == -1 || wasSecondNotEq
}
