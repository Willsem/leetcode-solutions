func largestPathValue(colors string, edges [][]int) int {
	n := len(colors)
	col := make([]int, n)
	for i := range col {
		col[i] = int(colors[i] - 'a')
	}
	outdeg := make([]int, n)
	indeg := make([]int, n)
	for _, e := range edges {
		outdeg[e[0]]++
		indeg[e[1]]++
	}
	adj := make([][]int, n)
	for i := range adj {
		adj[i] = make([]int, 0, outdeg[i])
	}
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
	}
	dp := make([][26]int, n)
	queue := make([]int, n)
	head, tail := 0, 0
	for i := 0; i < n; i++ {
		if indeg[i] == 0 {
			dp[i][col[i]] = 1
			queue[tail] = i
			tail++
		}
	}
	seen, ans := 0, 0
	for head < tail {
		u := queue[head]
		head++
		seen++
		for c := 0; c < 26; c++ {
			if dp[u][c] > ans {
				ans = dp[u][c]
			}
		}
		rowu := dp[u]
		for _, v := range adj[u] {
			rowv := &dp[v]
			cv := col[v]
			for c := 0; c < 26; c++ {
				val := rowu[c]
				if c == cv {
					val++
				}
				if val > (*rowv)[c] {
					(*rowv)[c] = val
				}
			}
			indeg[v]--
			if indeg[v] == 0 {
				queue[tail] = v
				tail++
			}
		}
		dp[u] = [26]int{}
	}
	if seen != n {
		return -1
	}
	return ans
}
