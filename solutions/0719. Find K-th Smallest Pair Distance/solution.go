import "sort"

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	return binarySearch(-1, nums[len(nums)-1]-nums[0], func(maxDist int) bool {
		count := 0
		l := 0
		for r := range nums {
			for nums[r]-nums[l] > maxDist {
				l++
			}
			count += r - l
		}
		return count < k
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
