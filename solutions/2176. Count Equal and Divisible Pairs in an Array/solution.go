func countPairs(nums []int, k int) int {
	count := 0
	index := make(map[int][]int)
	for i := range nums {
		for _, j := range index[nums[i]] {
			if i*j%k == 0 {
				count++
			}
		}

		index[nums[i]] = append(index[nums[i]], i)
	}

	return count
}
