func minCapability(nums []int, k int) int {
	maxNum := nums[0]
	for _, num := range nums {
		maxNum = max(maxNum, num)
	}

	return binarySearch(0, maxNum, func(mid int) bool {
		possibleThefts := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] <= mid {
				possibleThefts++
				i++
			}
		}

		return possibleThefts < k
	}, true)
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
