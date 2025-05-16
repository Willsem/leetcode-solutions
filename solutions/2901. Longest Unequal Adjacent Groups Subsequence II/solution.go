func getWordsInLongestSubsequence(words []string, groups []int) []string {
	dp := make([]int, len(words))
	prev := make([]int, len(words))
	for i := range dp {
		dp[i] = 1
		prev[i] = -1
	}
	maxIndex := 0

	for i := 1; i < len(dp); i++ {
		for j := 0; j < i; j++ {
			if check(words[i], words[j]) && dp[j]+1 > dp[i] && groups[i] != groups[j] {
				dp[i] = dp[j] + 1
				prev[i] = j
			}
		}
		if dp[i] > dp[maxIndex] {
			maxIndex = i
		}
	}

	ans := []string{}
	for i := maxIndex; i >= 0; i = prev[i] {
		ans = append(ans, words[i])
	}

	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-i-1] = ans[len(ans)-i-1], ans[i]
	}

	return ans
}

func check(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	diff := 0
	for i := range s1 {
		if s1[i] != s2[i] {
			diff++
			if diff > 1 {
				return false
			}
		}
	}
	return diff == 1
}
