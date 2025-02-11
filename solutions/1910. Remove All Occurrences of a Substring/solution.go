func removeOccurrences(s string, part string) string {
	result := make([]byte, 0)
	for i := range s {
		result = append(result, s[i])

		if len(result) >= len(part) {
			isEqual := true
			for j := range part {
				if result[len(result)-len(part)+j] != part[j] {
					isEqual = false
					break
				}
			}

			if isEqual {
				result = result[:len(result)-len(part)]
			}
		}
	}

	return string(result)
}
