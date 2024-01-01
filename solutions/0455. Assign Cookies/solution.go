import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	contentIndex, cookieIndex := 0, 0
	for cookieIndex < len(s) && contentIndex < len(g) {
		if s[cookieIndex] >= g[contentIndex] {
			contentIndex++
		}
		cookieIndex++
	}

	return contentIndex
}
