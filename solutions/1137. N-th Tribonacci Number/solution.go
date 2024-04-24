func tribonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	}

	prev1, prev2, prev3 := 0, 1, 1
	for i := 3; i <= n; i++ {
		prev1, prev2, prev3 = prev2, prev3, prev1+prev2+prev3
	}

	return prev3
}
