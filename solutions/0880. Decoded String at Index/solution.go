func decodeAtIndex(s string, k int) string {
	length := int64(0)
	i := 0

	for length < int64(k) {
		if s[i] >= '0' && s[i] <= '9' {
			length *= int64(s[i] - '0')
		} else {
			length++
		}
		i++
	}

	for j := i - 1; j >= 0; j-- {
		if s[j] >= '0' && s[j] <= '9' {
			length /= int64(s[j] - '0')
			k %= int(length)
		} else {
			if k == 0 || k == int(length) {
				return string(s[j])
			}
			length--
		}
	}

	return ""
}
