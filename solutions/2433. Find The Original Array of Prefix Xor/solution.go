func findArray(pref []int) []int {
	for i := len(pref) - 1; i > 0; i-- {
		pref[i] ^= pref[i-1]
	}

	return pref
}
