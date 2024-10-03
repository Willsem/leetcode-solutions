func minSubarray(nums []int, p int) int {
	sum := 0
	for _, num := range nums {
		sum += num
		sum %= p
	}

	target := sum % p
	if target == 0 {
		return 0
	}

	modMap := make(map[int]int)
	modMap[0] = -1

	currSum, minLen := 0, len(nums)

	for i := 0; i < len(nums); i++ {
		currSum += nums[i]
		currSum %= p

		need := (currSum - target + p) % p
		if found, ok := modMap[need]; ok {
			minLen = min(minLen, i-found)
		}

		modMap[currSum] = i
	}

	if minLen == len(nums) {
		return -1
	}

	return minLen
}
