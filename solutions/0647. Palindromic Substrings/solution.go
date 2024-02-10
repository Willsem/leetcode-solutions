func countSubstrings(s string) int {
	count := 0
	for i := range s {
		count += isPalindrome(s, i, i) + isPalindrome(s, i, i+1)
	}
	return count
}

func isPalindrome(s string, l, r int) int {
	res := 0
	for l >= 0 && r < len(s) && s[l] == s[r] {
		res++
		l--
		r++
	}
	return res
}
