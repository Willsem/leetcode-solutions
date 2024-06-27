func findCenter(edges [][]int) int {
	counts := make(map[int]int)
	center := 0
	for _, edge := range edges {
		counts[edge[0]]++
		counts[edge[1]]++

		if counts[center] < counts[edge[0]] {
			center = edge[0]
		}
		if counts[center] < counts[edge[1]] {
			center = edge[1]
		}
	}
	return center
}
