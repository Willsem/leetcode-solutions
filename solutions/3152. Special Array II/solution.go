func isArraySpecial(nums []int, queries [][]int) []bool {
	res := make([]bool, len(queries))

	prefix := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		prefix[i] = prefix[i-1]
		if nums[i]%2 == nums[i-1]%2 {
			prefix[i]++
		}
	}

	for i, q := range queries {
		start, end := q[0], q[1]
		res[i] = prefix[end]-prefix[start] == 0
	}

	return res
}
