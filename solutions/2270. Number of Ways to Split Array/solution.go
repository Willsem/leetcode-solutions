func waysToSplitArray(nums []int) int {
	n := len(nums)

	leftSum, rightSum := make([]int, n), make([]int, n)
	lsum, rsum := 0, 0
	for i := 0; i < n; i++ {
		lsum += nums[i]
		leftSum[i] = lsum

		rsum += nums[n-i-1]
		rightSum[n-i-1] = rsum
	}

	count := 0
	for i := 0; i < n-1; i++ {
		if leftSum[i] >= rightSum[i+1] {
			count++
		}
	}
	return count
}
