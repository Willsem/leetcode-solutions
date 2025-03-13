func minZeroArray(nums []int, queries [][]int) int {
	totalSum := 0
	k := 0
	diffArr := make([]int, len(nums)+1)

	for i := range nums {
		for totalSum+diffArr[i] < nums[i] {
			k++
			if k > len(queries) {
				return -1
			}

			q := queries[k-1]
			l, r, v := q[0], q[1], q[2]

			if r >= i {
				diffArr[max(l, i)] += v
				diffArr[r+1] -= v
			}
		}

		totalSum += diffArr[i]
	}

	return k
}
