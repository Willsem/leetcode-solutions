import (
	"sort"
	"strings"
)

type countChar struct {
	count int
	char  rune
}

func frequencySort(s string) string {
	counts := make(map[rune]countChar)
	for _, c := range s {
		if _, ok := counts[c]; !ok {
			counts[c] = countChar{
				count: 0,
				char:  c,
			}
		}

		count := counts[c]
		count.count++
		counts[c] = count
	}

	chars := make([]countChar, 0, len(counts))
	for _, count := range counts {
		chars = append(chars, count)
	}

	sort.Slice(chars, func(i, j int) bool {
		return chars[i].count < chars[j].count
	})

	var sb strings.Builder
	for _, count := range chars {
		for i := 0; i < count.count; i++ {
			sb.WriteString(string(count.char))
		}
	}
	return sb.String()
}
