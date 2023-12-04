func largestGoodInteger(num string) string {
	count, prev, max := 0, rune(0), rune(0)
	for _, c := range num {
		if prev == rune(0) {
			prev = c
			count = 1
			continue
		}

		if c == prev {
			count++

			if count >= 3 && c > max {
				max = c
			}
		} else {
			count = 1
		}

		prev = c
	}

	if max == rune(0) {
		return ""
	}

	return string(max) + string(max) + string(max)
}
