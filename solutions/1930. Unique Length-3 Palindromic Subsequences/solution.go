func countPalindromicSubsequence(s string) int {
	order := make([][2]int, 26)
	for i := range order {
		order[i] = [2]int{-1, -1}
	}

	for i, ch := range s {
		idx := ch - 'a'
		if order[idx][0] != -1 {
			order[idx][1] = i
		} else {
			order[idx][0] = i
		}
	}

	count := 0
	for _, pos := range order {
		if pos[1]-pos[0] > 1 {
			unique := make(map[rune]struct{})
			for _, midChar := range s[pos[0]+1 : pos[1]] {
				unique[midChar] = struct{}{}
			}
			count += len(unique)
		}
	}

	return count
}
