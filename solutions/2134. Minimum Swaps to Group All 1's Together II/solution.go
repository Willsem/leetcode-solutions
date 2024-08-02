func minSwaps(nums []int) int {
	n := len(nums)
	totalOnes := 0

	for _, num := range nums {
		totalOnes += num
	}

	if totalOnes == 0 || totalOnes == n {
		return 0
	}

	currentOnes := 0

	for i := 0; i < totalOnes; i++ {
		currentOnes += nums[i]
	}

	maxOnes := currentOnes

	for i := 0; i < n; i++ {
		currentOnes -= nums[i]
		currentOnes += nums[(i+totalOnes)%n]
		if currentOnes > maxOnes {
			maxOnes = currentOnes
		}
	}

	return totalOnes - maxOnes
}
