func checkValidString(s string) bool {
	open, close := 0, 0
	for i := range s {
		if s[i] == '(' || s[i] == '*' {
			open++
		} else {
			open--
		}

		if s[len(s)-i-1] == ')' || s[len(s)-i-1] == '*' {
			close++
		} else {
			close--
		}

		if open < 0 || close < 0 {
			return false
		}
	}
	return true
}
