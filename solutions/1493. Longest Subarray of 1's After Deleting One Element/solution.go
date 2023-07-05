func longestSubarray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	countOnes := 0
	seqLen := 0
	maxLen := 0

	for _, num := range nums {
		if num == 1 {
			countOnes++
			seqLen++

			if seqLen > maxLen {
				maxLen = seqLen
			}
		} else {
			seqLen = countOnes
			countOnes = 0
		}
	}

	if maxLen == len(nums) {
		maxLen--
	}

	return maxLen
}
