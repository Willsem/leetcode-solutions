func constructDistancedSequence(n int) []int {
	result := make([]int, 2*n-1)

	var dfs func(i, bitMap int) bool
	dfs = func(i, bitMap int) bool {
		if i == len(result) {
			return true
		}

		if result[i] > 0 {
			return dfs(i+1, bitMap)
		}

		for num := n; num > 0; num-- {
			if bitMap&(1<<num) != 0 {
				continue
			}

			if num > 1 && (i+num >= len(result) || result[i+num] > 0) {
				continue
			}

			result[i] = num
			if num > 1 {
				result[i+num] = num
			}

			if dfs(i+1, bitMap^(1<<num)) {
				return true
			}

			result[i] = 0
			if num > 1 {
				result[i+num] = 0
			}
		}

		return false
	}

	dfs(0, 0)
	return result
}
