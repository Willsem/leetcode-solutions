func triangleType(nums []int) string {
	switch {
	case nums[0] >= nums[1]+nums[2] || nums[1] >= nums[0]+nums[2] || nums[2] >= nums[0]+nums[1]:
		return "none"

	case nums[0] == nums[1] && nums[0] == nums[2]:
		return "equilateral"

	case nums[0] == nums[1] || nums[0] == nums[2] || nums[1] == nums[2]:
		return "isosceles"

	default:
		return "scalene"
	}
}
