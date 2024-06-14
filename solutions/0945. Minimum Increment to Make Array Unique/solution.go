func minIncrementForUnique(nums []int) int {
	max := nums[0]
	for _, num := range nums {
		if max < num {
			max = num
		}
	}

	counts := make([]int, len(nums)+max+1)
	for _, num := range nums {
		counts[num]++
	}

	res := 0
	for i, count := range counts {
		if count <= 1 {
			continue
		}

		duplicates := count - 1
		counts[i+1] += duplicates
		counts[i] = 1
		res += duplicates
	}
	return res
}
