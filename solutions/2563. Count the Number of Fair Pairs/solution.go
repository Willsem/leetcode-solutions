import "sort"

func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	return countPairs(nums, upper) - countPairs(nums, lower-1)
}

func countPairs(nums []int, target int) int64 {
	var count int64
	l, r := 0, len(nums)-1

	for l < r {
		if nums[l]+nums[r] > target {
			r--
		} else {
			count += int64(r - l)
			l++
		}
	}

	return count
}
