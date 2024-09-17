import "strings"

func uncommonFromSentences(s1 string, s2 string) []string {
	countAppears := make(map[string]int)
	countWords(countAppears, s1)
	countWords(countAppears, s2)

	result := make([]string, 0)
	for word, count := range countAppears {
		if count == 1 {
			result = append(result, word)
		}
	}
	return result
}

func countWords(countAppears map[string]int, s string) {
	for _, word := range strings.Split(s, " ") {
		countAppears[word]++
	}
}
