import "strings"

func makeFancyString(s string) string {
	res := &strings.Builder{}

	var (
		prevChar rune
		count    int
	)
	for _, v := range s {
		if v == prevChar {
			count++
		} else {
			count = 1
			prevChar = v
		}

		if count < 3 {
			res.WriteRune(v)
		}
	}

	return res.String()
}
