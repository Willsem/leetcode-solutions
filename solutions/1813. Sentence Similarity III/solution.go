import "strings"

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	words1 := strings.Split(sentence1, " ")
	words2 := strings.Split(sentence2, " ")

	if len(words1) < len(words2) {
		words1, words2 = words2, words1
	}

	start, end := 0, 0

	for start < len(words2) && words1[start] == words2[start] {
		start++
	}
	for end < len(words2) && words1[len(words1)-end-1] == words2[len(words2)-end-1] {
		end++
	}

	return start+end >= len(words2)
}
