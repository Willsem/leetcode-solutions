func largestVariance(s string) int {
	counter := make(map[rune]int, 26)
	for _, c := range s {
		if _, ok := counter[c]; !ok {
			counter[c] = 0
		}
		counter[c]++
	}

	globalMax := 0

	for i := 'a'; i <= 'z'; i++ {
		for j := 'a'; j <= 'z'; j++ {
			if i == j || counter[i] == 0 || counter[j] == 0 {
				continue
			}

			major := i
			minor := j
			majorCount := 0
			minorCount := 0

			restMinor := counter[j]

			for _, c := range s {
				if c == major {
					majorCount++
				}

				if c == minor {
					minorCount++
					restMinor--
				}

				if minorCount > 0 && globalMax < majorCount-minorCount {
					globalMax = majorCount - minorCount
				}

				if majorCount < minorCount && restMinor > 0 {
					majorCount = 0
					minorCount = 0
				}

			}
		}
	}

	return globalMax
}
