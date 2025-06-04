func lastSubstring(s string) string {
	i, j, n := 0, 1, len(s)
	for j < n {
		k := 0
		for j+k < n && s[i+k] == s[j+k] {
			k++
		}
		if j+k < n && s[i+k] < s[j+k] {
			i, j = j, max(j+1, i+k+1)
		} else {
			j = j + k + 1
		}
	}
	return s[i:]
}

func answerString(word string, numFriends int) string {
	if numFriends == 1 {
		return word
	}

	last := lastSubstring(word)
	n, m := len(word), len(last)
	return last[:min(m, n-numFriends+1)]
}
