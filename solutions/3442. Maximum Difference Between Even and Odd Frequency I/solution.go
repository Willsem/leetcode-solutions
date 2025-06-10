func maxDifference(s string) int {
	freq := make(map[rune]int)
	for _, c := range s {
		freq[c]++
	}

	maxOdd := -1
	minEven := -1
	for _, count := range freq {
		if count%2 == 0 && (minEven == -1 || minEven > count) {
			minEven = count
		}
		if count%2 != 0 && (maxOdd == -1 || maxOdd < count) {
			maxOdd = count
		}
	}

	return maxOdd - minEven
}
