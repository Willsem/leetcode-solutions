func largestAltitude(gain []int) int {
	sum := 0
	maxSum := 0
	for _, v := range gain {
		sum += v
		if maxSum < sum {
			maxSum = sum
		}
	}
	return maxSum
}
