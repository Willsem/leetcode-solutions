func isArraySpecial(nums []int) bool {
	prevParity := nums[0] % 2
	for _, num := range nums[1:] {
		if prevParity == num%2 {
			return false
		}
		prevParity = num % 2
	}
	return true
}
