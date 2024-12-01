func checkIfExist(arr []int) bool {
	m := make(map[int]struct{})
	for _, num := range arr {
		_, okMulti := m[num*2]
		_, okDivide := m[num/2]

		if okMulti || (num%2 == 0 && okDivide) {
			return true
		}

		m[num] = struct{}{}
	}

	return false
}
