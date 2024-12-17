import "strings"

func repeatLimitedString(s string, repeatLimit int) string {
	counts := make([]int, 26)

	for _, c := range s {
		counts[int(c-'a')]++
	}

	i, j := 25, 25
	sb := &strings.Builder{}

	for j >= 0 {
		if i == j || (i >= 0 && counts[i] == 0) {
			i--
			continue
		}

		if counts[j] == 0 {
			j--
			continue
		}

		count := min(repeatLimit, counts[j])
		counts[j] -= count
		sb.WriteString(strings.Repeat(string('a'+j), count))

		if i < 0 {
			break
		}

		if counts[j] > 0 {
			sb.WriteRune(rune('a' + i))
			counts[i]--
		} else {
			j = i
		}
	}

	return sb.String()
}
