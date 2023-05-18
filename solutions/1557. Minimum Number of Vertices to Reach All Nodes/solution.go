func findSmallestSetOfVertices(n int, edges [][]int) []int {
	degrees := make([]int, n)
	for _, edge := range edges {
		degrees[edge[1]]++
	}

	ans := make([]int, 0)
	for i := range degrees {
		if degrees[i] == 0 {
			ans = append(ans, i)
		}
	}

	return ans
}
