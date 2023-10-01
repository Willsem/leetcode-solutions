func reverseWords(s string) string {
	res := make([]byte, 0, len(s))
	word := make([]byte, 0)

	for _, v := range s {
		if v == ' ' {
			res = append(res, word...)
			res = append(res, ' ')
			word = word[:0]
		} else {
			word = append([]byte{byte(v)}, word...)
		}
	}

	if s[len(s)-1] != ' ' {
		res = append(res, word...)
	}

	return string(res)
}
