func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	for _, v := range prerequisites {
		a := v[0]
		b := v[1]

		graph[b] = append(graph[b], a)
	}

	inStack := make([]bool, numCourses)
	visited := make([]bool, numCourses)

	var dfs func(curr int) bool
	dfs = func(curr int) bool {
		visited[curr] = true
		inStack[curr] = true

		for _, node := range graph[curr] {
			if inStack[node] {
				return true
			}

			if !visited[node] && dfs(node) {
				return true
			}
		}

		inStack[curr] = false
		return false
	}

	for i := range visited {
		if !visited[i] && dfs(i) {
			return false
		}
	}

	return true
}
