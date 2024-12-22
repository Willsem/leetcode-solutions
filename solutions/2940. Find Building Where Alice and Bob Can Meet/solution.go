func leftmostBuildingQueries(heights []int, queries [][]int) []int {
	result := make([]int, len(queries))
	for i := range result {
		result[i] = -1
	}
	monoStack := make([][2]int, 0, 16)
	newQueries := make([][][2]int, len(heights))

	for i := 0; i < len(queries); i++ {
		a := queries[i][0]
		b := queries[i][1]
		if a > b {
			a, b = b, a
		}
		if heights[a] < heights[b] || a == b {
			result[i] = b
		} else {
			newQueries[b] = append(newQueries[b], [2]int{heights[a], i})
		}
	}

	for i := len(heights) - 1; i >= 0; i-- {
		ssize := len(monoStack)
		for _, q := range newQueries[i] {
			height := q[0]
			query := q[1]
			pos := search(height, monoStack)
			if pos < ssize && pos >= 0 {
				result[query] = monoStack[pos][1]
			}
		}
		for len(monoStack) > 0 && monoStack[len(monoStack)-1][0] <= heights[i] {
			monoStack = monoStack[:len(monoStack)-1]
		}
		monoStack = append(monoStack, [2]int{heights[i], i})
	}

	return result
}

func search(height int, monoStack [][2]int) int {
	l, r := 0, len(monoStack)-1
	ans := -1

	for l <= r {
		m := (l + (r-l)/2)
		if monoStack[m][0] > height {
			ans = max(ans, m)
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return ans
}
