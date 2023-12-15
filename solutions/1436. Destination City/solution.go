func destCity(paths [][]string) string {
	counts := make(map[string]int)

	for _, path := range paths {
		counts[path[0]]++
		if _, ok := counts[path[1]]; !ok {
			counts[path[1]] = 0
		}
	}

	for k, v := range counts {
		if v == 0 {
			return k
		}
	}

	return ""
}
