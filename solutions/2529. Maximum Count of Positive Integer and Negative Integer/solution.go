func maximumCount(nums []int) int {
	if nums[len(nums)-1] < 0 || nums[0] > 0 {
		return len(nums)
	}

	firstZeroPos := binarySearch(-1, len(nums)-1, func(i int) bool {
		return nums[i] < 0
	}, true)
	lastZeroPos := binarySearch(0, len(nums), func(i int) bool {
		return nums[i] <= 0
	}, false)

	neg := firstZeroPos
	pos := len(nums) - lastZeroPos - 1
	return max(neg, pos)
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
