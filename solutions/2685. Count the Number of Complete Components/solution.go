func countCompleteComponents(n int, edges [][]int) int {
	graph := make(map[int][]int)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	groupToNodes := make(map[int][]int)
	visited := make(map[int]struct{})

	var dfs func(node, group int)
	dfs = func(node, group int) {
		if _, ok := visited[node]; ok {
			return
		}

		visited[node] = struct{}{}
		groupToNodes[group] = append(groupToNodes[group], node)
		for _, next := range graph[node] {
			dfs(next, group)
		}
	}

	currGroup := 0
	for node := range n {
		if _, ok := visited[node]; !ok {
			dfs(node, currGroup)
			currGroup++
		}
	}

	result := 0
	for _, group := range groupToNodes {
		isComplete := true
		for _, node := range group {
			if len(graph[node]) < len(group)-1 {
				isComplete = false
				break
			}
		}
		if isComplete {
			result++
		}
	}
	return result
}
