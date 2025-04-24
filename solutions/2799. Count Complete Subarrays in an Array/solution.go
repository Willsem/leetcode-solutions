func countCompleteSubarrays(nums []int) int {
	distinct := make(map[int]bool)
	for _, num := range nums {
		distinct[num] = true
	}
	distinctCount := len(distinct)

	result, right := 0, 0
	count := make(map[int]int)
	for left := range nums {
		if left > 0 {
			count[nums[left-1]]--
			if count[nums[left-1]] == 0 {
				delete(count, nums[left-1])
			}
		}

		for right < len(nums) && len(count) < distinctCount {
			count[nums[right]]++
			right++
		}

		if len(count) == distinctCount {
			result += (len(nums) - right + 1)
		}
	}

	return result
}
