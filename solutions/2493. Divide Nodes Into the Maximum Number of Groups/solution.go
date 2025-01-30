func magnificentSets(n int, edges [][]int) int {
	graph := make([][]int, n)

	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	heights := make(map[int]int)

	for start := 0; start < n; start++ {
		if componentID, height, ok := bfsCheckAndFindHeight(start, graph); ok {
			if h, exists := heights[componentID]; exists {
				heights[componentID] = max(h, height)
			} else {
				heights[componentID] = height
			}
		} else {
			return -1
		}
	}

	totalDiameters := 0
	for _, height := range heights {
		totalDiameters += height
	}
	return totalDiameters
}

func bfsCheckAndFindHeight(start int, graph [][]int) (componentID int, maxHeight int, ok bool) {
	level := make([]int, len(graph))
	queue := []int{start}
	level[start] = 1
	maxHeight = 1
	componentID = start

	for front := 0; front < len(queue); front++ {
		v := queue[front]
		maxHeight = level[v]
		componentID = min(componentID, v)

		for _, u := range graph[v] {
			if level[u] == 0 {
				level[u] = level[v] + 1
				queue = append(queue, u)
			} else if abs(level[u]-level[v]) != 1 {
				ok = false
				return
			}
		}
	}

	ok = true
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
