import "sort"

func minDeletions(s string) int {
	counts := make(map[rune]int)
	for _, c := range s {
		counts[c]++
	}

	sortedFreqs := make([]int, 0, len(counts))
	for _, freq := range counts {
		sortedFreqs = append(sortedFreqs, freq)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sortedFreqs)))

	res := 0
	usedFreq := make(map[int]bool)
	for _, freq := range sortedFreqs {
		for freq > 0 && usedFreq[freq] {
			freq--
			res++
		}
		usedFreq[freq] = true
	}

	return res
}
