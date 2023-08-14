func findKthLargest(nums []int, k int) int {
	maxValue := nums[0]
	minValue := nums[0]

	for _, v := range nums {
		if v > maxValue {
			maxValue = v
		}
		if v < minValue {
			minValue = v
		}
	}

	counts := make([]int, maxValue-minValue+1)

	for _, v := range nums {
		counts[v-minValue] += 1
	}

	for i := len(counts) - 1; i >= 0; i-- {
		k -= counts[i]
		if k <= 0 {
			return i + minValue
		}
	}

	return -1
}
