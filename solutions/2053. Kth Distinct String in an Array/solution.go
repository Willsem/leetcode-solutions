func kthDistinct(arr []string, k int) string {
	counts := make(map[string]int)

	for _, str := range arr {
		counts[str]++
	}

	i := 0
	for _, str := range arr {
		if counts[str] == 1 {
			i++
			if i == k {
				return str
			}
		}
	}

	return ""
}
