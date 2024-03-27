func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}

	total, prod := 0, 1

	left := 0
	for right, num := range nums {
		prod *= num

		for prod >= k {
			prod /= nums[left]
			left++
		}

		total += right - left + 1
	}

	return total
}
