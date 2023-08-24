import "strings"

func fullJustify(words []string, maxWidth int) []string {
	res := make([]string, 0)
	curr := make([]string, 0)
	numOfLetters := 0

	for _, word := range words {
		if len(word)+len(curr)+numOfLetters > maxWidth {
			for i := 0; i < maxWidth-numOfLetters; i++ {
				curr[i%max(1, len(curr)-1)] += " "
			}
			res = append(res, strings.Join(curr, ""))
			curr = curr[:0]
			numOfLetters = 0
		}
		curr = append(curr, word)
		numOfLetters += len(word)
	}

	lastLine := strings.Join(curr, " ")
	for len(lastLine) < maxWidth {
		lastLine += " "
	}
	res = append(res, lastLine)

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
