func findDuplicate(nums []int) int {
	for _, num := range nums {
		n := abs(num)
		if nums[n] < 0 {
			return n
		}

		nums[n] *= -1
	}

	return -1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
