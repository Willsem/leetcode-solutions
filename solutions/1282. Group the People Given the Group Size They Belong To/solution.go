func groupThePeople(groupSizes []int) [][]int {
	ans := make([][]int, 0)
	groups := make(map[int][]int, 0)

	for i, size := range groupSizes {
		groups[size] = append(groups[size], i)

		if len(groups[size]) == size {
			ans = append(ans, groups[size])
			groups[size] = []int{}
		}
	}

	return ans
}
