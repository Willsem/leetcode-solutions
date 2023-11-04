func getLastMoment(n int, left []int, right []int) int {
	result := 0

	for _, num := range left {
		if result < num {
			result = num
		}
	}

	for _, num := range right {
		if result < n-num {
			result = n - num
		}
	}

	return result
}
