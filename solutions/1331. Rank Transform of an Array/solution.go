import "sort"

func arrayRankTransform(arr []int) []int {
	sorted := make([]int, len(arr))
	copy(sorted, arr)
	sort.Ints(sorted)

	mapping := make(map[int]int)
	rank := 1
	for _, num := range sorted {
		if _, ok := mapping[num]; !ok {
			mapping[num] = rank
			rank++
		}
	}

	for i := range arr {
		arr[i] = mapping[arr[i]]
	}
	return arr
}
