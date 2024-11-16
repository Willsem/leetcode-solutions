func resultsArray(nums []int, k int) []int {
	res := make([]int, 0, len(nums))
	powerLen := 1
	prevNum := -1
	for i, num := range nums {
		if prevNum+1 == num {
			powerLen++
		} else {
			powerLen = 1
		}
		prevNum = num

		if i < k-1 {
			continue
		}

		if powerLen >= k {
			res = append(res, num)
		} else {
			res = append(res, -1)
		}
	}

	return res
}
