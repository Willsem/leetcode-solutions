func maxDistance(arrays [][]int) int {
	maxVal, minVal := arrays[0][len(arrays[0])-1], arrays[0][0]
	ans := 0

	for _, arr := range arrays[1:] {
		maxDist := max(maxVal-arr[0], arr[len(arr)-1]-minVal)
		if maxDist > ans {
			ans = maxDist
		}

		minVal = min(minVal, arr[0])
		maxVal = max(maxVal, arr[len(arr)-1])
	}

	return ans
}
