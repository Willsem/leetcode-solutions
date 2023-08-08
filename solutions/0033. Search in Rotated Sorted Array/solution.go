func search(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1

	for l <= r {
		mid := (l + r) / 2
		if nums[mid] > nums[n-1] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	binarySearch := func(l, r, target int) int {
		for l <= r {
			mid := (l + r) / 2
			if nums[mid] == target {
				return mid
			} else if nums[mid] < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}

		return -1
	}

	if ans := binarySearch(0, l-1, target); ans != -1 {
		return ans
	}

	return binarySearch(l, n-1, target)
}
