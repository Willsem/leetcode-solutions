func minSubArrayLen(target int, nums []int) int {
	start, end := 0, 0
	sum := nums[0]
	min := -1

	for {
		if sum < target {
			end += 1

			if end == len(nums) {
				break
			}

			sum += nums[end]
		} else {
			if min == -1 || end-start+1 < min {
				min = end - start + 1
			}

			sum -= nums[start]
			start += 1
		}
	}

	if min == -1 {
		return 0
	}

	return min
}
