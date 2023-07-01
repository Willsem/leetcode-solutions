import "sort"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(cookies []int, k int, index int, children []int, unfairness int) int {
	if index < 0 {
		return unfairness
	}
	minUnfairness := 1<<31 - 1
	for i := 0; i < k; i++ {
		if max(children[i]+cookies[index], unfairness) >= minUnfairness {
			continue
		}
		children[i] += cookies[index]
		minUnfairness = min(minUnfairness, dfs(cookies, k, index-1, children, max(children[i], unfairness)))
		children[i] -= cookies[index]
	}
	return minUnfairness
}

func distributeCookies(cookies []int, k int) int {
	sort.Slice(cookies, func(i, j int) bool {
		return cookies[i] > cookies[j]
	})
	children := make([]int, k)
	return dfs(cookies, k, len(cookies)-1, children, 0)
}
