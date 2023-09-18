import "sort"

type strongInfo struct {
	index, strong int
}

func kWeakestRows(mat [][]int, k int) []int {
	strong := make([]strongInfo, len(mat))
	for i := range strong {
		curr := binarySearch(-1, len(mat[i]), func(j int) bool {
			return mat[i][j] == 1
		})

		strong[i] = strongInfo{
			index:  i,
			strong: curr + 1,
		}
	}

	sort.Slice(strong, func(i, j int) bool {
		if strong[i].strong == strong[j].strong {
			return strong[i].index < strong[j].index
		}

		return strong[i].strong < strong[j].strong
	})

	result := make([]int, k)
	for i := range result {
		result[i] = strong[i].index
	}

	return result
}

func binarySearch(l, r int, comp func(i int) bool) int {
	for l < r-1 {
		mid := (l + r) / 2
		if comp(mid) {
			l = mid
		} else {
			r = mid
		}
	}

	return l
}
