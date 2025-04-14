func countGoodTriplets(arr []int, a int, b int, c int) int {
	count := 0
	for i := range arr {
		for j := i + 1; j < len(arr); j++ {
			for k := j + 1; k < len(arr); k++ {
				aDiff := abs(arr[i] - arr[j])
				bDiff := abs(arr[j] - arr[k])
				cDiff := abs(arr[i] - arr[k])

				if aDiff <= a && bDiff <= b && cDiff <= c {
					count++
				}
			}
		}
	}

	return count
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
