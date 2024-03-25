func findDuplicates(nums []int) []int {
	ans := make([]int, 0)
	for _, num := range nums {
		n := int(math.Abs(float64(num)))
		if nums[n-1] < 0 {
			ans = append(ans, n)
		} else {
			nums[n-1] *= -1
		}
	}
	return ans
}
