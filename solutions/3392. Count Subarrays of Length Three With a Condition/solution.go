func countSubarrays(nums []int) int {
	count := 0
	for i := 0; i < len(nums)-2; i++ {
		if (nums[i]+nums[i+2])*2 == nums[i+1] {
			count++
		}
	}
	return count
}
