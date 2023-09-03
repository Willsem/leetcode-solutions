func uniquePaths(m int, n int) int {
	ans := 1
	for i := 1; i < m; i++ {
		ans = ans * (n - 1 + i) / i
	}
	return ans
}
