func countGood(nums []int, k int) int64 {
	l, count, result := 0, 0, 0
	counts := make(map[int]int)

	for r := range nums {
		count += counts[nums[r]]
		counts[nums[r]]++

		for l < len(nums) && count-counts[nums[l]]+1 >= k {
			counts[nums[l]]--
			count -= counts[nums[l]]
			l++
		}

		if count >= k {
			result += l + 1
		}
	}

	return int64(result)
}
