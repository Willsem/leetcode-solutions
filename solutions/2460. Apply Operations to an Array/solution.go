func applyOperations(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] *= 2
			nums[i+1] = 0
		}
	}

	zeroCount := 0
	for i := range nums {
		if nums[i] == 0 {
			zeroCount++
		} else {
			nums[i-zeroCount] = nums[i]
		}
	}
	for i := len(nums) - zeroCount; i < len(nums); i++ {
		nums[i] = 0
	}

	return nums
}
