func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	curr, prev := 2, 1
	for i := 3; i <= n; i++ {
		curr, prev = curr+prev, curr
	}

	return curr
}
