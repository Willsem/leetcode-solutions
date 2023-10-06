func integerBreak(n int) int {
	if n <= 3 {
		return n - 1
	}

	memo := make([]int, n+1)

	var dp func(num int) int
	dp = func(num int) int {
		if num <= 3 {
			return num
		}

		if memo[num] != 0 {
			return memo[num]
		}

		ans := num
		for i := 2; i < num; i++ {
			curr := i * dp(num-i)
			if curr > ans {
				ans = curr
			}
		}

		memo[num] = ans
		return ans
	}

	return dp(n)
}
