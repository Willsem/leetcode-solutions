func maximumBeauty(nums []int, k int) int {
	if len(nums) == 1 {
		return 1
	}

	maxValue := nums[0]
	for _, num := range nums {
		maxValue = max(maxValue, num)
	}

	count := make([]int, maxValue+1)

	for _, num := range nums {
		count[max(num-k, 0)]++
		if num+k+1 <= maxValue {
			count[num+k+1]--
		}
	}

	maxBeaty, currSum := 0, 0
	for _, v := range count {
		currSum += v
		maxBeaty = max(maxBeaty, currSum)
	}

	return maxBeaty
}
