import "sort"

func heightChecker(heights []int) int {
	sorted := make([]int, len(heights))
	copy(sorted, heights)
	sort.Ints(sorted)

	count := 0
	for i := range heights {
		if heights[i] != sorted[i] {
			count++
		}
	}
	return count
}
