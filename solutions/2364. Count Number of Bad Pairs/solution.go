func countBadPairs(nums []int) int64 {
	count := 0
	diffCounts := make(map[int]int)

	for i := range nums {
		diff := i - nums[i]
		diffCount := diffCounts[diff]
		count += i - diffCount
		diffCounts[diff] = diffCount + 1
	}

	return int64(count)
}
