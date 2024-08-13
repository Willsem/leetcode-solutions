import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	sort.Ints(candidates)

	var dfs func(path []int, target int, currI int)
	dfs = func(path []int, target int, currI int) {
		if target < 0 {
			return
		}

		if target == 0 {
			newPath := make([]int, len(path))
			copy(newPath, path)
			ans = append(ans, newPath)
			return
		}

		for i := currI; i < len(candidates); i++ {
			if i > currI && candidates[i] == candidates[i-1] {
				continue
			}

			path = append(path, candidates[i])
			dfs(path, target-candidates[i], i+1)
			path = path[:len(path)-1]
		}
	}

	dfs([]int{}, target, 0)
	return ans
}
