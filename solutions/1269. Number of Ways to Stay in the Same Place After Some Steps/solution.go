const mod = 1e9 + 7

func numWays(steps int, arrLen int) int {
	memo := make(map[int]map[int]int)

	var dp func(curr, remain int) int
	dp = func(curr, remain int) int {
		if remain == 0 {
			if curr == 0 {
				return 1
			}

			return 0
		}

		if currMemo, ok := memo[curr]; ok {
			if res, ok := currMemo[remain]; ok {
				return res
			}
		}

		ans := dp(curr, remain-1) % mod
		if curr > 0 {
			ans = (ans + dp(curr-1, remain-1)) % mod
		}

		if curr < arrLen-1 {
			ans = (ans + dp(curr+1, remain-1)) % mod
		}

		if _, ok := memo[curr]; !ok {
			memo[curr] = make(map[int]int)
		}

		memo[curr][remain] = ans
		return ans
	}

	return dp(0, steps)
}
