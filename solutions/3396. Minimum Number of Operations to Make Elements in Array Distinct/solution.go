func minimumOperations(nums []int) int {
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}

	countNonDistinct := 0
	for _, count := range counts {
		if count > 1 {
			countNonDistinct++
		}
	}

	operations := 0
	i := 0
	for countNonDistinct > 0 && i < len(nums) {
		for range 3 {
			if i >= len(nums) {
				continue
			}

			counts[nums[i]]--
			if counts[nums[i]] == 1 {
				countNonDistinct--
			}
			i++
		}

		operations++
	}

	return operations
}
