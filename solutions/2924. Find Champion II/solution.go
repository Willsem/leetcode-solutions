func findChampion(n int, edges [][]int) int {
	indegree := make([]int, n)
	for _, edge := range edges {
		indegree[edge[1]]++
	}

	count := 0
	champion := -1
	for i := range indegree {
		if indegree[i] == 0 {
			count++
			champion = i
		}
	}

	if count > 1 {
		return -1
	}
	return champion
}
