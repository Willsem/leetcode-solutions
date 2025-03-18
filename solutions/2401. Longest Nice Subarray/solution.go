func longestNiceSubarray(nums []int) int {
	maxLen, bits := 0, 0
	left := 0
	for right := range nums {
		for bits&nums[right] != 0 {
			bits ^= nums[left]
			left++
		}

		bits |= nums[right]
		maxLen = max(maxLen, right-left+1)
	}

	return maxLen
}
