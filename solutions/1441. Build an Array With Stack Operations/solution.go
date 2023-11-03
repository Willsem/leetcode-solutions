func buildArray(target []int, n int) []string {
	curr := 0
	stream := 1
	result := make([]string, 0)

	for curr < len(target) {
		result = append(result, "Push")
		if stream == target[curr] {
			curr++
		} else {
			result = append(result, "Pop")
		}

		stream++
	}

	return result
}
