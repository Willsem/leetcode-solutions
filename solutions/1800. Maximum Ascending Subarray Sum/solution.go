func maxAscendingSum(nums []int) int {
	maxSum, currSum := nums[0], nums[0]
	prev := nums[0]

	for _, num := range nums[1:] {
		if num > prev {
			currSum += num
		} else {
			currSum = num
		}

		maxSum = max(maxSum, currSum)
		prev = num
	}

	return maxSum
}
