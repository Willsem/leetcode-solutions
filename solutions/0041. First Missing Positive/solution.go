func firstMissingPositive(nums []int) int {
	has1 := false
	for i := range nums {
		if nums[i] == 1 {
			has1 = true
		} else if nums[i] <= 0 || nums[i] > len(nums) {
			nums[i] = 1
		}
	}

	if !has1 {
		return 1
	}

	for _, num := range nums {
		n := int(math.Abs(float64(num)))
		if n == len(nums) {
			nums[0] = -int(math.Abs(float64(nums[0])))
		} else {
			nums[n] = -int(math.Abs(float64(nums[n])))
		}
	}

	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			return i
		}
	}

	if nums[0] > 0 {
		return len(nums)
	}

	return len(nums) + 1
}
