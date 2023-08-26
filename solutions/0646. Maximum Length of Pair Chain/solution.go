import "sort"

func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})

	curr, ans := pairs[0][1], 1
	for _, pair := range pairs {
		if pair[0] > curr {
			ans++
			curr = pair[1]
		}
	}

	return ans
}
