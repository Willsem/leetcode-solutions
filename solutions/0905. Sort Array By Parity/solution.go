import "sort"

func sortArrayByParity(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i]%2 == 0
	})
	return nums
}
