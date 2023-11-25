func getSumAbsoluteDifferences(nums []int) []int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	prevSum := 0
	for i, num := range nums {
		sum -= num
		nums[i] = (num*i - prevSum) + (sum - num*(len(nums)-i-1))
		prevSum += num
	}

	return nums
}
