import "sort"

func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)
	left := 0
	curr := int64(0)

	for right := 0; right < len(nums); right++ {
		target := nums[right]
		curr += int64(target)

		if int64((right-left+1)*target)-curr > int64(k) {
			curr -= int64(nums[left])
			left++
		}
	}

	return len(nums) - left
}
