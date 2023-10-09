func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	left := binarySearch(-1, len(nums)-1, func(i int) bool {
		return nums[i] < target
	}, true)
	if nums[left] != target {
		return []int{-1, -1}
	}

	right := binarySearch(0, len(nums), func(i int) bool {
		return nums[i] <= target
	}, false)

	return []int{left, right}
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
