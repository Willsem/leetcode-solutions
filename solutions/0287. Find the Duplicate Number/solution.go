func findDuplicate(nums []int) int {
	for _, num := range nums {
		n := int(math.Abs(float64(num)))
		if nums[n-1] < 0 {
			return n
		}
		nums[n-1] *= -1
	}
	return -1
}
