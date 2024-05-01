import "strings"

func reversePrefix(word string, ch byte) string {
	pos := 0
	for i := range word {
		if word[i] == ch {
			pos = i
			break
		}
	}

	sb := strings.Builder{}
	for i := pos; i >= 0; i-- {
		sb.WriteByte(word[i])
	}

	sb.WriteString(word[pos+1:])

	return sb.String()
}
