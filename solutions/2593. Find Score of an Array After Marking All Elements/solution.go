func findScore(nums []int) int64 {
	ans := 0
	for i := 0; i < len(nums); i += 2 {
		n := i
		for i+1 < len(nums) && nums[i] > nums[i+1] {
			i++
		}
		for j := i; j >= n; j -= 2 {
			ans += nums[j]
		}
	}
	return int64(ans)
}
