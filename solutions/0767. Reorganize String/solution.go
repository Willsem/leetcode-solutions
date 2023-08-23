import "sort"

func reorganizeString(s string) string {
	freqMap := make(map[rune]int)
	for _, c := range s {
		freqMap[c]++
	}

	var sortedChars []rune
	for ch := range freqMap {
		sortedChars = append(sortedChars, ch)
	}

	sort.Slice(sortedChars, func(i, j int) bool {
		return freqMap[sortedChars[i]] > freqMap[sortedChars[j]]
	})

	if freqMap[sortedChars[0]] > (len(s)+1)/2 {
		return ""
	}

	res := make([]rune, len(s))
	i := 0
	for _, ch := range sortedChars {
		for j := 0; j < freqMap[ch]; j++ {
			if i >= len(s) {
				i = 1
			}
			res[i] = ch
			i += 2
		}
	}

	return string(res)
}
