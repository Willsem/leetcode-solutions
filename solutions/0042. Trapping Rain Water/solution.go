func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	left, right := make([]int, len(height)), make([]int, len(height))

	left[0] = height[0]
	for i := 1; i < len(left); i++ {
		left[i] = max(left[i-1], height[i])
	}

	right[len(right)-1] = height[len(height)-1]
	for i := len(right) - 2; i >= 0; i-- {
		right[i] = max(right[i+1], height[i])
	}

	result := 0
	for i := range height {
		result += min(left[i], right[i]) - height[i]
	}
	return result
}
