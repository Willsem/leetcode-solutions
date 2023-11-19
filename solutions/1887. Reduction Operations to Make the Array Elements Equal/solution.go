import "sort"

func reductionOperations(nums []int) int {
	sort.Ints(nums)

	ans, up := 0, 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			up++
		}

		ans += up
	}

	return ans
}
