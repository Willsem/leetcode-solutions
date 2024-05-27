func specialArray(nums []int) int {
	counts := make(map[int]int, len(nums))
	maxValue := nums[0]
	for _, num := range nums {
		counts[num]++
		maxValue = max(maxValue, num)
	}

	prefixSum := 0
	for i := maxValue; i >= 0; i-- {
		prefixSum += counts[i]
		if i == prefixSum {
			return i
		}
	}

	return -1
}
