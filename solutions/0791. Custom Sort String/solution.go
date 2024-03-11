import "strings"

func customSortString(order string, s string) string {
	counts := make(map[byte]int)
	for _, c := range s {
		counts[byte(c)]++
	}

	res := strings.Builder{}
	was := make(map[byte]struct{})
	for _, c := range order {
		if count, ok := counts[byte(c)]; ok {
			delete(counts, byte(c))
			was[byte(c)] = struct{}{}
			for i := 0; i < count; i++ {
				res.WriteString(string(c))
			}
		}
	}

	for _, c := range s {
		if _, ok := was[byte(c)]; !ok {
			res.WriteString(string(c))
		}
	}

	return res.String()
}
