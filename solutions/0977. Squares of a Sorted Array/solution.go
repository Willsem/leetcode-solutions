func sortedSquares(nums []int) []int {
	ans := make([]int, len(nums))

	left, right, curr := 0, len(nums)-1, len(nums)-1
	for left <= right {
		lSquare := nums[left] * nums[left]
		rSquare := nums[right] * nums[right]

		if lSquare > rSquare {
			ans[curr] = lSquare
			left++
		} else {
			ans[curr] = rSquare
			right--
		}
		curr--
	}

	return ans
}
