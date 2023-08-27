func canCross(stones []int) bool {
	if stones[1] != 1 {
		return false
	}

	dp := map[[2]int]bool{}
	var dfs func(i int, unit int) bool
	dfs = func(i int, unit int) bool {
		if i == len(stones)-1 {
			return true
		}

		if v, f := dp[[2]int{i, unit}]; f {
			return v
		}

		res := false
		for j := i + 1; j < len(stones); j++ {
			newUnit := stones[j] - stones[i]
			if newUnit >= unit-1 && newUnit <= unit+1 && dfs(j, newUnit) {
				res = true
				break
			}
		}
		dp[[2]int{i, unit}] = res
		return dp[[2]int{i, unit}]
	}
	return dfs(1, 1)
}
