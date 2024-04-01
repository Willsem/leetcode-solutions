func lengthOfLastWord(s string) int {
	p := len(s) - 1
	for p > 0 && s[p] == ' ' {
		p--
	}

	start := p
	for p >= 0 && s[p] != ' ' {
		p--
	}

	return start - p
}
