import "strings"

func addSpaces(s string, spaces []int) string {
	result := &strings.Builder{}

	prev := 0
	for _, i := range spaces {
		result.WriteString(s[prev:i])
		result.WriteString(" ")
		prev = i
	}

	result.WriteString(s[prev:])
	return result.String()
}
