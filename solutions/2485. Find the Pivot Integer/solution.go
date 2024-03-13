func pivotInteger(n int) int {
	sum := n * (n + 1) / 2
	currSum := 0
	for i := 1; i <= n; i++ {
		currSum += i
		if currSum == sum+i-currSum {
			return i
		}
	}
	return -1
}
