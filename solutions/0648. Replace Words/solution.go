import "strings"

func replaceWords(dictionary []string, sentence string) string {
	roots := make(map[string]struct{})
	for _, root := range dictionary {
		roots[root] = struct{}{}
	}

	result := &strings.Builder{}
	for _, word := range strings.Split(sentence, " ") {
		result.WriteString(" ")

		was := false
		for i := 0; i < len(word); i++ {
			if _, ok := roots[word[:i]]; ok {
				result.WriteString(word[:i])
				was = true
				break
			}
		}

		if !was {
			result.WriteString(word)
		}
	}

	return result.String()[1:]
}
