func minPatches(nums []int, n int) int {
	miss := 1
	result := 0

	for i := 0; miss <= n; {
		if i < len(nums) && nums[i] <= miss {
			miss += nums[i]
			i++
		} else {
			miss += miss
			result++
		}
	}

	return result
}
