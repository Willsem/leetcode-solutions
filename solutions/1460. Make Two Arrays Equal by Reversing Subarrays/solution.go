func canBeEqual(target []int, arr []int) bool {
	mapTarget := make(map[int]int)
	mapArr := make(map[int]int)

	for _, num := range target {
		mapTarget[num]++
	}
	for _, num := range arr {
		mapArr[num]++
	}

	if len(mapTarget) != len(mapArr) {
		return false
	}

	for key := range mapTarget {
		if mapTarget[key] != mapArr[key] {
			return false
		}
	}

	return true
}
