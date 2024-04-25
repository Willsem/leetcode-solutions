func longestIdealString(s string, k int) int {
	dp := make([]int, len(s))
	lastLetter := make(map[byte]int)
	res := 0

	for i := range s {
		choise := 1

		for diff := -k; diff <= k; diff++ {
			if j, ok := lastLetter[s[i]-byte(diff)]; ok {
				choise = max(choise, dp[j]+1)
			}
		}

		dp[i] = choise
		lastLetter[s[i]] = i
		res = max(res, dp[i])
	}

	return res
}
