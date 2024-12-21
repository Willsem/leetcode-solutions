func maxKDivisibleComponents(n int, edges [][]int, values []int, k int) int {
	graph := make(map[int]map[int]bool)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		if graph[u] == nil {
			graph[u] = make(map[int]bool)
		}
		if graph[v] == nil {
			graph[v] = make(map[int]bool)
		}
		graph[u][v] = true
		graph[v][u] = true
	}

	res := 1

	var dfs func(node int) int
	dfs = func(node int) int {
		total := values[node]
		for nextNode := range graph[node] {
			delete(graph[nextNode], node)
			nextTotal := dfs(nextNode)

			if nextTotal%k == 0 {
				res++
			} else {
				total += nextTotal
			}
		}

		return total
	}

	dfs(0)
	return res
}
