func minimumSubarrayLength(nums []int, k int) int {
	minLen := len(nums) + 1

	bitsCount := make([]int, 32)
	left := 0
	for right := 0; right < len(nums); right++ {
		updateBitsCount(bitsCount, nums[right], 1)

		for left <= right && countBitsToNumber(bitsCount) >= k {
			minLen = min(minLen, right-left+1)
			updateBitsCount(bitsCount, nums[left], -1)
			left++
		}
	}

	if minLen == len(nums)+1 {
		return -1
	}

	return minLen
}

func updateBitsCount(bitsCount []int, num int, delta int) {
	for bit := range bitsCount {
		if (num>>bit)&1 == 1 {
			bitsCount[bit] += delta
		}
	}
}

func countBitsToNumber(bitsCount []int) int {
	num := 0
	for bit := range bitsCount {
		if bitsCount[bit] != 0 {
			num |= 1 << bit
		}
	}
	return num
}
