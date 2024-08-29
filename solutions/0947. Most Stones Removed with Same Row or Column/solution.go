func removeStones(stones [][]int) int {
	count := 0
	set := make(map[int]int)

	var find func(int) int
	find = func(el int) int {
		if _, ok := set[el]; !ok {
			set[el] = el
			count++
		}
		if set[el] != el {
			set[el] = find(set[el])
		}
		return set[el]
	}

	merge := func(elementA, elementB int) {
		repA := find(elementA)
		repB := find(elementB)
		if repA != repB {
			set[repB] = repA
			count--
		}
	}

	for _, stone := range stones {
		merge(stone[0]+1, stone[1]+10002)
	}

	return len(stones) - count
}
