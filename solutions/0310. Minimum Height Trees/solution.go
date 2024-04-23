func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	graph := make(map[int][]int, n)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	leaves := make([]int, 0)
	for i := range graph {
		if len(graph[i]) == 1 {
			leaves = append(leaves, i)
		}
	}

	for n > 2 {
		nextLeaves := make([]int, 0)

		for _, leaf := range leaves {
			neighbor := graph[leaf][0]
			for i, node := range graph[neighbor] {
				if node == leaf {
					graph[neighbor] = append(graph[neighbor][:i], graph[neighbor][i+1:]...)
					break
				}
			}

			if len(graph[neighbor]) == 1 {
				nextLeaves = append(nextLeaves, neighbor)
			}
			n--
		}

		leaves = nextLeaves
	}

	return leaves
}
