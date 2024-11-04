import (
	"strconv"
	"strings"
)

func compressedString(word string) string {
	result := &strings.Builder{}

	var (
		prev  rune
		count int
	)
	for _, c := range word {
		if prev == c {
			count++
		} else {
			writeToResult(result, count, prev)
			prev = c
			count = 1
		}

		if count == 9 {
			writeToResult(result, count, c)
			prev, count = 0, 0
		}
	}

	writeToResult(result, count, prev)
	return result.String()
}

func writeToResult(result *strings.Builder, count int, c rune) {
	if count == 0 {
		return
	}

	countStr := strconv.Itoa(count)
	result.WriteString(countStr)
	result.WriteRune(c)
}
