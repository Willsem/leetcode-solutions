import "strings"

func divideString(s string, k int, fill byte) []string {
	builder := &strings.Builder{}
	result := make([]string, 0, len(s)/k+1)

	for _, c := range s {
		builder.WriteRune(c)

		if builder.Len() == k {
			result = append(result, builder.String())
			builder.Reset()
		}
	}

	if builder.Len() > 0 {
		for builder.Len() != k {
			builder.WriteByte(fill)
		}
		result = append(result, builder.String())
	}

	return result
}
