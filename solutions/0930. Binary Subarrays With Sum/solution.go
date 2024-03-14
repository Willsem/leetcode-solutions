func numSubarraysWithSum(nums []int, goal int) int {
	currSum, prefixZeros := 0, 0
	left, right := 0, 0
	count := 0
	for right < len(nums) {
		currSum += nums[right]

		for left < right && (nums[left] == 0 || currSum > goal) {
			if nums[left] == 1 {
				prefixZeros = 0
			} else {
				prefixZeros++
			}

			currSum -= nums[left]
			left++
		}

		if currSum == goal {
			count += 1 + prefixZeros
		}

		right++
	}

	return count
}
