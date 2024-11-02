import "strings"

func isCircularSentence(sentence string) bool {
	words := strings.Split(sentence, " ")

	firstWord, lastWord := words[0], words[len(words)-1]
	if firstWord[0] != lastWord[len(lastWord)-1] {
		return false
	}

	prev := firstWord
	for i := 1; i < len(words); i++ {
		word := words[i]
		if prev[len(prev)-1] != word[0] {
			return false
		}
		prev = word
	}

	return true
}
