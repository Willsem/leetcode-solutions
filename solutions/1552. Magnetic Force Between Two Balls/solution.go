import "sort"

func maxDistance(position []int, m int) int {
	sort.Ints(position)

	ans := -1

	binarySearch(0, position[len(position)-1]-position[0]+1, func(mid int) bool {
		lastPosition, balls := position[0], 1
		for i := 1; i < len(position); i++ {
			if position[i]-lastPosition >= mid {
				lastPosition = position[i]
				balls++
			}
		}

		if balls >= m {
			ans = mid
		}

		return balls >= m
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
