func shortestPalindrome(s string) string {
	count := kmp(reverse(s), s)
	return reverse(s[count:]) + s
}

func kmp(txt, patt string) int {
	newString := patt + "#" + txt
	pi := make([]int, len(newString))
	i, k := 1, 0
	for i < len(newString) {
		if newString[i] == newString[k] {
			k++
			pi[i] = k
			i++
		} else {
			if k > 0 {
				k = pi[k-1]
			} else {
				pi[i] = 0
				i++
			}
		}
	}
	return pi[len(newString)-1]
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
