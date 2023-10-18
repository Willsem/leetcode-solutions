func minimumTime(n int, relations [][]int, time []int) int {
	graph := make(map[int][]int)
	in_degree := make([]int, n+1)
	for _, relation := range relations {
		graph[relation[0]] = append(graph[relation[0]], relation[1])
		in_degree[relation[1]]++
	}

	dist := make([]int, n+1)
	copy(dist[1:], time)
	queue := []int{}
	for i := 1; i <= n; i++ {
		if in_degree[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		for _, next_course := range graph[course] {
			dist[next_course] = max(dist[next_course], dist[course]+time[next_course-1])
			in_degree[next_course]--
			if in_degree[next_course] == 0 {
				queue = append(queue, next_course)
			}
		}
	}

	maxVal := 0
	for _, val := range dist {
		maxVal = max(maxVal, val)
	}
	return maxVal
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
