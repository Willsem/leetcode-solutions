func isMonotonic(nums []int) bool {
	inc := true
	dec := true

	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			inc = false
		}
		if nums[i] > nums[i-1] {
			dec = false
		}
		if !inc && !dec {
			return false
		}
	}

	return true
}
