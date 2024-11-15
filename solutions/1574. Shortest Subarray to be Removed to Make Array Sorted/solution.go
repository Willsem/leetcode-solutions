func findLengthOfShortestSubarray(arr []int) int {
	r := len(arr) - 1
	for r > 0 && arr[r] >= arr[r-1] {
		r--
	}

	ans := r
	for l := 0; l < r && (l == 0 || arr[l-1] <= arr[l]); l++ {
		for r < len(arr) && arr[l] > arr[r] {
			r++
		}
		ans = min(ans, r-l-1)
	}

	return ans
}
