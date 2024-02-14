func rearrangeArray(nums []int) []int {
	res := make([]int, len(nums))

	pPos, pNeg := 0, 1
	for _, num := range nums {
		if num >= 0 {
			res[pPos] = num
			pPos += 2
		} else {
			res[pNeg] = num
			pNeg += 2
		}
	}

	return res
}
