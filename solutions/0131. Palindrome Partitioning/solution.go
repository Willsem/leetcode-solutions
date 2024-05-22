func partition(s string) [][]string {
	res := make([][]string, 0)
	var curr []string

	var backtrack func(idx int)
	backtrack = func(idx int) {
		if idx == len(s) {
			res = append(res, append([]string{}, curr...))
			return
		}

		for i := idx; i < len(s); i++ {
			subStr := s[idx : i+1]
			if isPalindrome(subStr) {
				curr = append(curr, subStr)
				backtrack(i + 1)
				curr = curr[:len(curr)-1]
			}
		}
	}

	backtrack(0)
	return res
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
