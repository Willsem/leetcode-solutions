import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	priority := make(map[int]int, len(arr2))
	for i, num := range arr2 {
		priority[num] = i
	}

	sort.Slice(arr1, func(i, j int) bool {
		pi, oki := priority[arr1[i]]
		pj, okj := priority[arr1[j]]

		switch {
		case oki && okj:
			return pi < pj
		case oki:
			return true
		case okj:
			return false
		default:
			return arr1[i] < arr1[j]
		}
	})

	return arr1
}
