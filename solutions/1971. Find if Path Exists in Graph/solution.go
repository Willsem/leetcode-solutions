func validPath(n int, edges [][]int, source int, destination int) bool {
	graph := make(map[int][]int)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	visited := make(map[int]struct{})

	var dfs func(node int) bool
	dfs = func(node int) bool {
		if node == destination {
			return true
		}

		if _, ok := visited[node]; ok {
			return false
		}

		visited[node] = struct{}{}

		for _, next := range graph[node] {
			if dfs(next) {
				return true
			}
		}

		return false
	}

	return dfs(source)
}
