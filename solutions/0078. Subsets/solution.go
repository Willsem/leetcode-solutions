func subsets(nums []int) [][]int {
	sets := make([][]int, 0)

	var collect func(set []int, i int)
	collect = func(set []int, i int) {
		if i == len(nums) {
			copySet := make([]int, len(set))
			copy(copySet, set)
			sets = append(sets, copySet)
			return
		}

		collect(set, i+1)
		set = append(set, nums[i])
		collect(set, i+1)
		set = set[:len(set)-1]
	}

	collect([]int{}, 0)
	return sets
}
