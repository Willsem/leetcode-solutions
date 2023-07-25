func peakIndexInMountainArray(arr []int) int {
	l := 0
	r := len(arr) - 1
	for r-l > 1 {
		mid := (l + r) / 2
		if arr[mid] > arr[mid-1] {
			l = mid
		} else {
			r = mid
		}
	}
	return l
}
