func maximumDifference(nums []int) int {
	maxDiff, currMin := -1, nums[0]
	for _, num := range nums {
		if num <= currMin {
			currMin = num
		} else {
			maxDiff = max(maxDiff, num-currMin)
		}
	}
	return maxDiff
}
