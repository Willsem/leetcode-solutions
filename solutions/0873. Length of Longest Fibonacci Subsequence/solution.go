type IntPair struct {
	first  int
	second int
}

func lenLongestFibSubseq(arr []int) int {
	maxLen := 0
	dp := make(map[IntPair]int)
	seen := make(map[int]struct{})

	for k, z := range arr {
		seen[z] = struct{}{}

		for j := k - 1; j >= 0; j-- {
			y := arr[j]
			diff := z - y

			if diff >= y {
				break
			}

			if _, found := seen[diff]; found {
				curLen := max(dp[IntPair{y, diff}], 2) + 1
				maxLen = max(maxLen, curLen)

				dp[IntPair{z, y}] = curLen
			}
		}
	}

	return maxLen
}
