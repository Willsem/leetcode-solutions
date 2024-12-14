import "math"

func continuousSubarrays(nums []int) int64 {
	r, l := 0, 0
	total := 0

	currMin, currMax := nums[0], nums[0]

	for ; r < len(nums); r++ {
		currMin = min(currMin, nums[r])
		currMax = max(currMax, nums[r])

		if currMax-currMin > 2 {
			windowLen := r - l
			total += (windowLen * (windowLen + 1)) / 2

			l = r
			currMin = nums[r]
			currMax = nums[r]

			for l > 0 && int(math.Abs(float64(nums[r]-nums[l-1]))) <= 2 {
				l--
				currMin = min(currMin, nums[l])
				currMax = max(currMax, nums[l])
			}

			windowLen = r - l
			total -= (windowLen * (windowLen + 1)) / 2
		}
	}

	windowLen := r - l
	total += (windowLen * (windowLen + 1)) / 2

	return int64(total)
}
