type orderEl struct {
	l int
	r int
}

func countPalindromicSubsequence(s string) int {
	order := make([]orderEl, 26)
	for i := range order {
		order[i] = orderEl{-1, -1}
	}

	for i, ch := range s {
		idx := ch - 'a'
		if order[idx].l != -1 {
			order[idx].r = i
		} else {
			order[idx].l = i
		}
	}

	count := 0
	for _, pos := range order {
		if pos.r-pos.l > 1 {
			unique := make(map[rune]struct{})
			for _, midChar := range s[pos.l+1 : pos.r] {
				unique[midChar] = struct{}{}
			}
			count += len(unique)
		}
	}

	return count
}
