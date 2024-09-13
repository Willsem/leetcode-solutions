func xorQueries(arr []int, queries [][]int) []int {
	res := make([]int, 0, len(queries))

	for i := 1; i < len(arr); i++ {
		arr[i] ^= arr[i-1]
	}

	for _, el := range queries {
		left, right := el[0], el[1]
		if left > 0 {
			res = append(res, arr[left-1]^arr[right])
		} else {
			res = append(res, arr[right])
		}
	}

	return res
}
