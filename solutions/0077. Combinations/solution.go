func combine(n int, k int) [][]int {
	result := make([][]int, 0)

	var solve func(buffer map[int]struct{}, depth int, lastEl int)
	solve = func(buffer map[int]struct{}, depth int, lastEl int) {
		if depth == k {
			curr := make([]int, 0, k)
			for key := range buffer {
				curr = append(curr, key)
			}
			result = append(result, curr)
			return
		}

		for i := lastEl + 1; i <= n; i++ {
			if _, ok := buffer[i]; !ok {
				buffer[i] = struct{}{}
				solve(buffer, depth+1, i)
				delete(buffer, i)
			}
		}
	}
	buffer := make(map[int]struct{}, k)
	solve(buffer, 0, 0)

	return result
}
