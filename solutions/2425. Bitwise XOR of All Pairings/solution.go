func xorAllNums(nums1 []int, nums2 []int) int {
	counts := make(map[int]int)

	for _, num := range nums1 {
		counts[num] += len(nums2)
	}
	for _, num := range nums2 {
		counts[num] += len(nums1)
	}

	result := 0
	for num, count := range counts {
		if count%2 == 1 {
			result ^= num
		}
	}
	return result
}
