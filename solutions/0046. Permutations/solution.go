func permute(nums []int) [][]int {
	n := len(nums)

	result := make([][]int, 0)

	var solve func(path map[int]struct{}, seq []int)
	solve = func(path map[int]struct{}, seq []int) {
		if len(path) == n {
			curr := make([]int, 0, len(seq))
			for _, v := range seq {
				curr = append(curr, nums[v])
			}
			result = append(result, curr)
			return
		}

		for i := 0; i < n; i++ {
			if _, ok := path[i]; !ok {
				path[i] = struct{}{}
				seq = append(seq, i)

				solve(path, seq)

				delete(path, i)
				seq = seq[:len(seq)-1]
			}
		}
	}

	solve(make(map[int]struct{}, 0), make([]int, 0))

	return result
}
