import "sort"

func minPairSum(nums []int) int {
	sort.Ints(nums)

	max := 0
	for i := 0; i < len(nums)/2; i++ {
		diff := nums[i] + nums[len(nums)-i-1]
		if diff > max {
			max = diff
		}
	}

	return max
}
