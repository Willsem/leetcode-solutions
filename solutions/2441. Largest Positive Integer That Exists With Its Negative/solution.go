func findMaxK(nums []int) int {
	existsNegative := make(map[int]struct{})
	for _, num := range nums {
		if num < 0 {
			existsNegative[-num] = struct{}{}
		}
	}

	ans := -1
	for _, num := range nums {
		if num > ans {
			if _, ok := existsNegative[num]; ok {
				ans = num
			}
		}
	}

	return ans
}
