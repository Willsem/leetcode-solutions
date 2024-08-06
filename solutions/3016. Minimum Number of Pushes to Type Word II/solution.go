import "sort"

func minimumPushes(word string) int {
	counts := make([]int, 26)
	for _, l := range word {
		counts[int(l)-int('a')]++
	}

	sort.Sort(sort.Reverse(sort.IntSlice(counts)))

	res := 0
	for i := range counts {
		res += (i/8 + 1) * counts[i]
	}

	return res
}
