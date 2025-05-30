func closestMeetingNode(edges []int, node1 int, node2 int) int {
	dist1 := make([]int, len(edges))
	dist2 := make([]int, len(edges))

	for i := range edges {
		dist1[i] = -1
		dist2[i] = -1
	}

	dfs(node1, dist1, edges)
	dfs(node2, dist2, edges)

	ans := -1
	for node := range edges {
		if min(dist1[node], dist2[node]) < 0 {
			continue
		}

		if ans == -1 || max(dist1[node], dist2[node]) < max(dist1[ans], dist2[ans]) {
			ans = node
		}
	}
	return ans
}

func dfs(node int, dist []int, edges []int) {
	currDist := 0
	for node != -1 && dist[node] == -1 {
		dist[node] = currDist
		currDist++
		node = edges[node]
	}
}
