func getMaximumXor(nums []int, maximumBit int) []int {
	xorProduct := 0

	for _, num := range nums {
		xorProduct ^= num
	}

	ans := make([]int, len(nums))
	mask := (1 << maximumBit) - 1

	for i := 0; i < len(nums); i++ {
		ans[i] = xorProduct ^ mask
		xorProduct ^= nums[len(nums)-i-1]
	}

	return ans
}
