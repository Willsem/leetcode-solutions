func minDays(bloomDay []int, m int, k int) int {
	lastDay := 0
	for _, day := range bloomDay {
		if day > lastDay {
			lastDay = day
		}
	}

	ans := -1
	binarySearch(0, lastDay+1, func(day int) bool {
		count, bouquets := 0, 0
		for i := range bloomDay {
			if bloomDay[i] <= day {
				count++

				if count == k {
					bouquets++
					count = 0
				}
			} else {
				count = 0
			}
		}

		if bouquets >= m {
			ans = day
		}

		return bouquets < m
	}, false)

	return ans
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
