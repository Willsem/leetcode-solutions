func findMaxLength(nums []int) int {
	m := make(map[int]int)
	m[0] = -1
	maxLen, count := 0, 0

	for i, num := range nums {
		if num == 1 {
			count++
		} else {
			count--
		}

		if c, ok := m[count]; ok {
			maxLen = max(maxLen, i-c)
		} else {
			m[count] = i
		}
	}

	return maxLen
}
