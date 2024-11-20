func takeCharacters(s string, k int) int {
	count := make([]int, 3)
	for _, c := range s {
		count[int(c)-int('a')]++
	}

	for _, c := range count {
		if c < k {
			return -1
		}
	}

	window := make([]int, 3)
	l, maxWindow := 0, 0

	for r := 0; r < len(s); r++ {
		window[int(s[r])-int('a')]++

		for l <= r && (count[0]-window[0] < k || count[1]-window[1] < k || count[2]-window[2] < k) {
			window[int(s[l])-int('a')]--
			l++
		}

		maxWindow = max(maxWindow, r-l+1)
	}

	return len(s) - maxWindow
}
