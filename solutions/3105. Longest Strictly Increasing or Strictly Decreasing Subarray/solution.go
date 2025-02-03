func longestMonotonicSubarray(nums []int) int {
	maxLen := 1
	currInc, currDec := 1, 1

	prev := nums[0]
	for _, num := range nums[1:] {
		if num > prev {
			currInc++
		} else {
			currInc = 1
		}

		if num < prev {
			currDec++
		} else {
			currDec = 1
		}

		maxLen = max(maxLen, max(currInc, currDec))
		prev = num
	}

	return maxLen
}
