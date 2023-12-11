func findSpecialInteger(arr []int) int {
	n := len(arr)
	candidates := []int{arr[n/4], arr[n/2], arr[3*n/4]}
	target := n / 4

	for _, candidate := range candidates {
		l := binarySearch(-1, n-1, func(i int) bool {
			return arr[i] < candidate
		}, true)
		r := binarySearch(0, n, func(i int) bool {
			return arr[i] <= candidate
		}, false)

		if r-l+1 > target {
			return candidate
		}
	}

	return -1
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
