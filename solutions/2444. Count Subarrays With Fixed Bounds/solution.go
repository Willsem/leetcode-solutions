func countSubarrays(nums []int, minK int, maxK int) int64 {
	ans := 0
	pmin, pmax := -1, -1
	left := 0

	for right := range nums {
		if nums[right] < minK || nums[right] > maxK {
			left = right + 1
			pmin = -1
			pmax = -1
		} else {
			if nums[right] == minK {
				pmin = right
			}
			if nums[right] == maxK {
				pmax = right
			}
			ans += max(0, min(pmin, pmax)-left+1)
		}
	}

	return int64(ans)
}
