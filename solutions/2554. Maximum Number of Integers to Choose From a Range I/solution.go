func maxCount(banned []int, n int, maxSum int) int {
	banMap := make(map[int]struct{})
	for _, banNum := range banned {
		banMap[banNum] = struct{}{}
	}

	sum, count := 0, 0
	for i := 1; i <= n; i++ {
		if _, ok := banMap[i]; ok {
			continue
		}

		sum += i
		if sum > maxSum {
			break
		}
		count++
	}

	return count
}
