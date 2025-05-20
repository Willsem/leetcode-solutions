func isZeroArray(nums []int, queries [][]int) bool {
	deltaArr := make([]int, len(nums)+1)
	for _, q := range queries {
		l, r := q[0], q[1]
		deltaArr[l]++
		deltaArr[r+1]--
	}

	currDelta := 0
	for i := range nums {
		currDelta += deltaArr[i]
		if currDelta < nums[i] {
			return false
		}
	}

	return true
}
