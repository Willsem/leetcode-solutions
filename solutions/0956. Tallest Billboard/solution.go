func tallestBillboard(rods []int) int {
	dp := make(map[int]int, 1)
	dp[0] = 0

	for _, r := range rods {
		newDp := copyDp(dp)

		for diff, taller := range dp {
			shorter := taller - diff
			if _, ok := newDp[diff+r]; !ok {
				newDp[diff+r] = taller + r
			} else {
				newDp[diff+r] = max(newDp[diff+r], taller+r)
			}

			newDiff := abs(shorter + r - taller)
			newTaller := max(shorter+r, taller)

			if _, ok := newDp[newDiff]; !ok {
				newDp[newDiff] = newTaller
			} else {
				newDp[newDiff] = max(newDp[newDiff], newTaller)
			}
		}

		dp = newDp
	}

	return dp[0]
}

func copyDp(dp map[int]int) map[int]int {
	newDp := make(map[int]int, len(dp))
	for key, value := range dp {
		newDp[key] = value
	}
	return newDp
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
