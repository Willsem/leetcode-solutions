import "sort"

func minimizeMax(nums []int, p int) int {
	sort.Ints(nums)
	n := len(nums)

	return binarySearch(0, nums[n-1]-nums[0], func(i int) bool {
		return countValidPairs(nums, i) >= p
	})
}

func countValidPairs(nums []int, threshold int) int {
	index, count := 0, 0

	for index < len(nums)-1 {
		if nums[index+1]-nums[index] <= threshold {
			count++
			index++
		}

		index++
	}

	return count
}

func binarySearch(l, r int, comp func(i int) bool) int {
	for l < r {
		mid := (l + r) / 2
		if comp(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}
