func maxAbsoluteSum(nums []int) int {
	minPrefixSum, maxPrefixSum := 0, 0
	prefixSum := 0

	for _, num := range nums {
		prefixSum += num
		minPrefixSum = min(minPrefixSum, prefixSum)
		maxPrefixSum = max(maxPrefixSum, prefixSum)
	}

	return maxPrefixSum - minPrefixSum
}
