func check(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	wasSwitched := false
	if nums[0] < nums[len(nums)-1] {
		wasSwitched = true
	}

	prev := nums[0]
	for _, num := range nums[1:] {
		if num < prev {
			if wasSwitched {
				return false
			}
			wasSwitched = true
		}
		prev = num
	}

	return true
}
