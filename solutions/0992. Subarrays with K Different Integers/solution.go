func subarraysWithKDistinct(nums []int, k int) int {
	return subarraysWithKDistinctAtMost(nums, k) - subarraysWithKDistinctAtMost(nums, k-1)
}

func subarraysWithKDistinctAtMost(nums []int, k int) int {
	counts := make(map[int]int)
	left := 0
	ans := 0

	for right := range nums {
		counts[nums[right]]++

		for len(counts) > k {
			counts[nums[left]]--
			if counts[nums[left]] == 0 {
				delete(counts, nums[left])
			}
			left++
		}

		ans += right - left + 1
	}

	return ans
}
