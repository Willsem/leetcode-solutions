func minimumMountainRemovals(nums []int) int {
	longestInc := make([]int, len(nums))
	longestDec := make([]int, len(nums))

	for i := range nums {
		longestInc[i] = 1
		longestDec[i] = 1
	}

	for i := range nums {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				longestInc[i] = max(longestInc[i], longestInc[j]+1)
			}
		}
	}

	for i := len(nums) - 1; i >= 0; i-- {
		for j := len(nums) - 1; j > i; j-- {
			if nums[i] > nums[j] {
				longestDec[i] = max(longestDec[i], longestDec[j]+1)
			}
		}
	}

	maxMountainLen := 0
	for i := 1; i < len(nums)-1; i++ {
		if longestInc[i] > 1 && longestDec[i] > 1 {
			maxMountainLen = max(maxMountainLen, longestInc[i]+longestDec[i]-1)
		}
	}

	return len(nums) - maxMountainLen
}
