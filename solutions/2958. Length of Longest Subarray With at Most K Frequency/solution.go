func maxSubarrayLength(nums []int, k int) int {
	freq := make(map[int]int)
	left := 0
	maxLen := 0

	for right := range nums {
		freq[nums[right]]++
		if freq[nums[right]] > k {
			for left < right {
				num := nums[left]
				freq[num]--
				left++

				if num == nums[right] {
					break
				}
			}
		}

		maxLen = max(maxLen, right-left+1)
	}

	return maxLen
}
