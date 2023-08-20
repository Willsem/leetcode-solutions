func sortItems(n int, m int, group []int, beforeItems [][]int) []int {
	// Isolated item is a group by itself
	for i := 0; i < n; i++ {
		if group[i] == -1 {
			group[i] = m
			m++
		}
	}

	adjGroup := make([][]int, m)
	adjItem := make([][]int, n)
	degGroup := make([]int, m)
	degItem := make([]int, n)

	for i := 0; i < m; i++ {
		adjGroup[i] = []int{}
	}
	for i := 0; i < n; i++ {
		adjItem[i] = []int{}
	}

	for i := 0; i < n; i++ {
		toI := group[i]
		for _, from := range beforeItems[i] {
			fromJ := group[from]
			if toI != fromJ {
				adjGroup[fromJ] = append(adjGroup[fromJ], toI)
				degGroup[toI]++
			}
			adjItem[from] = append(adjItem[from], i)
			degItem[i]++
		}
	}

	sortGroup := topoSort(adjGroup, degGroup, m)
	sortItem := topoSort(adjItem, degItem, n)
	if len(sortGroup) == 0 || len(sortItem) == 0 {
		return []int{}
	}

	itemGroup := make([][]int, m)
	for i := range itemGroup {
		itemGroup[i] = []int{}
	}
	for _, i := range sortItem {
		itemGroup[group[i]] = append(itemGroup[group[i]], i)
	}

	var ans []int
	for _, i := range sortGroup {
		ans = append(ans, itemGroup[i]...)
	}

	return ans
}

func topoSort(adj [][]int, deg []int, sz int) []int {
	q := []int{}
	for i := 0; i < sz; i++ {
		if deg[i] == 0 {
			q = append(q, i)
		}
	}
	res := []int{}
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		res = append(res, x)
		for _, y := range adj[x] {
			deg[y]--
			if deg[y] == 0 {
				q = append(q, y)
			}
		}
	}
	for i := 0; i < sz; i++ {
		if deg[i] > 0 {
			return []int{}
		}
	}
	return res
}
