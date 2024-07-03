import "sort"

func minDifference(nums []int) int {
	if len(nums) <= 4 {
		return 0
	}

	sort.Ints(nums)

	left, right := 0, len(nums)-4
	res := nums[right] - nums[left]
	for left < 4 {
		res = min(res, nums[right]-nums[left])
		left++
		right++
	}
	return res
}
