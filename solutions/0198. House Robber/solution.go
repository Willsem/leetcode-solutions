import "math"

func rob(nums []int) int {
	for i := 1; i < len(nums); i++ {
		if i == 1 {
			nums[i] = int(math.Max(float64(nums[i-1]), float64(nums[i])))
		} else {
			nums[i] = int(math.Max(float64(nums[i-1]), float64(nums[i-2]+nums[i])))
		}
	}

	return nums[len(nums)-1]
}
