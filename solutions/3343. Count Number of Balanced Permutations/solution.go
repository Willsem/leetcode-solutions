const mod = 1e9 + 7

func countBalancedPermutations(num string) int {
	tot, n := 0, len(num)
	count := make([]int, 10)
	for _, ch := range num {
		d := int(ch - '0')
		count[d]++
		tot += d
	}
	if tot%2 != 0 {
		return 0
	}

	target := tot / 2
	maxOdd := (n + 1) / 2
	comb := make([][]int, maxOdd+1)
	for i := range comb {
		comb[i] = make([]int, maxOdd+1)
		comb[i][i], comb[i][0] = 1, 1
		for j := 1; j < i; j++ {
			comb[i][j] = (comb[i-1][j] + comb[i-1][j-1]) % mod
		}
	}

	psum := make([]int, 11)
	for i := 9; i >= 0; i-- {
		psum[i] = psum[i+1] + count[i]
	}

	memo := make([][][]int, 10)
	for i := range memo {
		memo[i] = make([][]int, target+1)
		for j := range memo[i] {
			memo[i][j] = make([]int, maxOdd+1)
			for k := range memo[i][j] {
				memo[i][j][k] = -1
			}
		}
	}

	var dfs func(pos, curr, oddCount int) int
	dfs = func(pos, curr, oddCount int) int {
		if oddCount < 0 || psum[pos] < oddCount || curr > target {
			return 0
		}
		if pos > 9 {
			if curr == target && oddCount == 0 {
				return 1
			}
			return 0
		}
		if memo[pos][curr][oddCount] != -1 {
			return memo[pos][curr][oddCount]
		}
		evenCount := psum[pos] - oddCount
		res := 0
		start := max(0, count[pos]-evenCount)
		end := min(count[pos], oddCount)
		for i := start; i <= end; i++ {
			ways := comb[oddCount][i] * comb[evenCount][count[pos]-i] % mod
			res = (res + ways*dfs(pos+1, curr+i*pos, oddCount-i)%mod) % mod
		}
		memo[pos][curr][oddCount] = res
		return res
	}

	return dfs(0, 0, maxOdd)
}
