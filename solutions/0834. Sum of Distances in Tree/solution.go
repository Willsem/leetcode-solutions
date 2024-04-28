func sumOfDistancesInTree(n int, edges [][]int) []int {
	graph := make(map[int][]int)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	count := make([]int, n)
	res := make([]int, n)

	var dfs1 func(cur, parent int)
	dfs1 = func(cur, parent int) {
		count[cur] = 1
		for _, child := range graph[cur] {
			if child != parent {
				dfs1(child, cur)
				count[cur] += count[child]
				res[cur] += res[child] + count[child]
			}
		}
	}

	var dfs2 func(cur, parent int)
	dfs2 = func(cur, parent int) {
		for _, child := range graph[cur] {
			if child != parent {
				res[child] = res[cur] + n - 2*count[child]
				dfs2(child, cur)
			}
		}
	}

	dfs1(0, -1)
	dfs2(0, -1)

	return res
}
