func minimumDeletions(s string) int {
	n := len(s)
	minDel := 0
	countB := 0

	for i := 0; i < n; i++ {
		if s[i] == 'b' {
			countB++
		} else {
			minDel = min(minDel+1, countB)
		}
	}

	return minDel
}
