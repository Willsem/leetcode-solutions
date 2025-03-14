func maximumCandies(candies []int, k int64) int {
	return binarySearch(0, 10_000_001, func(mid int) bool {
		childrenCount := 0
		for _, candy := range candies {
			childrenCount += candy / mid
		}

		return int64(childrenCount) >= k
	}, false)
}

func binarySearch(l, r int, comp func(i int) bool, leftSearch bool) int {
	for l < r-1 {
		mid := (l + r) / 2
		if comp(mid) {
			l = mid
		} else {
			r = mid
		}
	}

	if leftSearch {
		return r
	}

	return l
}
