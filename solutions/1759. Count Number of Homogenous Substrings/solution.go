const mod = 1e9 + 7

func countHomogenous(s string) int {
	streak := 0
	ans := 0

	for i := range s {
		if i == 0 || s[i] == s[i-1] {
			streak++
		} else {
			streak = 1
		}

		ans = (ans + streak) % mod
	}

	return ans
}
