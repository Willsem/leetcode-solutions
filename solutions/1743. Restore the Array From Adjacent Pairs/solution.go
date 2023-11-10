func restoreArray(adjacentPairs [][]int) []int {
	graph := make(map[int][]int)

	for _, pair := range adjacentPairs {
		addToGraph(&graph, pair[0], pair[1])
		addToGraph(&graph, pair[1], pair[0])
	}

	startNode := -1
	for k, v := range graph {
		if len(v) == 1 {
			startNode = k
			break
		}
	}

	result := make([]int, 0, len(graph))

	var dfs func(node, prev int)
	dfs = func(node, prev int) {
		next := graph[node]

		if len(next) == 1 {
			if prev != -1 {
				return
			}

			result = append(result, next[0])
			dfs(next[0], node)
		} else {
			if next[0] == prev {
				result = append(result, next[1])
				dfs(next[1], node)
			} else {
				result = append(result, next[0])
				dfs(next[0], node)
			}
		}
	}

	result = append(result, startNode)
	dfs(startNode, -1)

	return result
}

func addToGraph(graph *map[int][]int, a, b int) {
	if _, ok := (*graph)[a]; !ok {
		(*graph)[a] = []int{b}
	} else {
		(*graph)[a] = append((*graph)[a], b)
	}
}
