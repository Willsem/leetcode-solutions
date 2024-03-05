func minimumLength(s string) int {
	left, right := 0, len(s)-1
	for left < right && s[left] == s[right] {
		curr := s[left]

		for left <= right && s[left] == curr {
			left++
		}

		for right > left && s[right] == curr {
			right--
		}
	}

	if left > right {
		return 0
	}

	return right - left + 1
}
