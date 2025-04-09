func minOperations(nums []int, k int) int {
	for _, num := range nums {
		if num < k {
			return -1
		}
	}

	diffNums := make(map[int]struct{})
	for _, num := range nums {
		if num > k {
			diffNums[num] = struct{}{}
		}
	}

	return len(diffNums)
}
