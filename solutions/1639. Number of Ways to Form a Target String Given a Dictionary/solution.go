const mod = 1e9 + 7

func numWays(words []string, target string) int {
	m, n := len(words[0]), len(target)

	freq := make([][]int, m)
	for i := range freq {
		freq[i] = make([]int, 26)
		for _, word := range words {
			freq[i][word[i]-'a']++
		}
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var solve func(targetPos int, colPos int) int
	solve = func(targetPos int, colPos int) int {
		if targetPos == len(target) {
			return 1
		}
		if colPos == m {
			return 0
		}

		if dp[targetPos][colPos] != -1 {
			return dp[targetPos][colPos]
		}

		result := 0

		count := freq[colPos][target[targetPos]-'a']
		if count > 0 {
			result = (count * solve(targetPos+1, colPos+1)) % mod
		}
		result = (result + solve(targetPos, colPos+1)) % mod

		dp[targetPos][colPos] = result
		return result
	}

	return solve(0, 0)
}
