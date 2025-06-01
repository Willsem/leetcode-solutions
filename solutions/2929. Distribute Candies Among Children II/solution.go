func distributeCandies(n int, limit int) int64 {
	result := 0

	for i := range min(limit, n) + 1 {
		if n-i > 2*limit {
			continue
		}

		result += min(n-i, limit) - max(0, n-i-limit) + 1
	}

	return int64(result)
}
