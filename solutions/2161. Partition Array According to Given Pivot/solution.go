func pivotArray(nums []int, pivot int) []int {
	result := make([]int, len(nums))
	l, r := 0, len(nums)-1

	for i, j := 0, len(nums)-1; i < len(nums); i, j = i+1, j-1 {
		if nums[i] < pivot {
			result[l] = nums[i]
			l++
		}

		if nums[j] > pivot {
			result[r] = nums[j]
			r--
		}
	}

	for l <= r {
		result[l] = pivot
		l++
	}

	return result
}
