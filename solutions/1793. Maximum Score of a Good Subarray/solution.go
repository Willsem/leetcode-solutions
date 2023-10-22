func maximumScore(nums []int, k int) int {
	left, right := k, k
	minVal := nums[k]
	maxScore := minVal

	for left > 0 || right < len(nums)-1 {
		if left == 0 || (right < len(nums)-1 && nums[right+1] > nums[left-1]) {
			right++
		} else {
			left--
		}
		minVal = min(minVal, min(nums[left], nums[right]))
		maxScore = max(maxScore, minVal*(right-left+1))
	}

	return maxScore
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
