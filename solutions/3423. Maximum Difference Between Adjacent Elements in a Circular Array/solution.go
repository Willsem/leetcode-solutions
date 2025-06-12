import "math"

func maxAdjacentDistance(nums []int) int {
	maxDiff := abs(nums[0] - nums[len(nums)-1])
	for i := 0; i < len(nums)-1; i++ {
		maxDiff = max(maxDiff, abs(nums[i]-nums[i+1]))
	}

	return maxDiff
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}
