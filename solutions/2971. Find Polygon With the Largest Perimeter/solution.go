import "sort"

func largestPerimeter(nums []int) int64 {
	sort.Ints(nums)
	prevElSum := 0
	ans := -1

	for _, num := range nums {
		if num < prevElSum {
			ans = num + prevElSum
		}
		prevElSum += num
	}

	return int64(ans)
}
