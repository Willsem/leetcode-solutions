func maxWidthRamp(nums []int) int {
	n := len(nums)

	rightMax := make([]int, n)
	rightMax[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(nums[i], rightMax[i+1])
	}

	l, r := 0, 0
	maxWidth := 0
	for r < n {
		for l < r && nums[l] > rightMax[r] {
			l++
		}
		maxWidth = max(maxWidth, r-l)
		r++
	}

	return maxWidth
}
