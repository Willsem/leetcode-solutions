func longestSubarray(nums []int) int {
	res, maxVal, currStreak := 0, 0, 0
	for _, num := range nums {
		if maxVal < num {
			maxVal = num
			res, currStreak = 0, 0
		}

		if maxVal == num {
			currStreak++
		} else {
			currStreak = 0
		}

		res = max(res, currStreak)
	}

	return res
}
