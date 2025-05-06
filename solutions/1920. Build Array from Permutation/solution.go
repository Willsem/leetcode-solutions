func buildArray(nums []int) []int {
	result := make([]int, len(nums))
	for i := range result {
		result[i] = nums[nums[i]]
	}
	return result
}
