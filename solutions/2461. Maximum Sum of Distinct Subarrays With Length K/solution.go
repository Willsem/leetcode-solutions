func maximumSubarraySum(nums []int, k int) int64 {
	sum, l, r := 0, 0, k-1
	counts := make(map[int]int, 0)
	repeats := make(map[int]struct{}, 0)
	for i := l; i <= r; i++ {
		sum += nums[i]
		counts[nums[i]]++

		if counts[nums[i]] > 1 {
			repeats[nums[i]] = struct{}{}
		}
	}

	maxSum := 0
	for r < len(nums) {
		if len(repeats) == 0 {
			maxSum = max(maxSum, sum)
		}

		sum -= nums[l]
		counts[nums[l]]--
		if counts[nums[l]] <= 1 {
			delete(repeats, nums[l])
		}
		l++

		r++
		if r < len(nums) {
			counts[nums[r]]++
			if counts[nums[r]] > 1 {
				repeats[nums[r]] = struct{}{}
			}
			sum += nums[r]
		}
	}

	return int64(maxSum)
}
