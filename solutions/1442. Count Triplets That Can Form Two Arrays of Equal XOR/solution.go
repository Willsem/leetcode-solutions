func countTriplets(arr []int) int {
	prefixXOR := make([]int, len(arr)+1)
	for i := range arr {
		prefixXOR[i+1] = prefixXOR[i] ^ arr[i]
	}

	count := 0
	for start := range prefixXOR {
		for end := start + 1; end < len(prefixXOR); end++ {
			if prefixXOR[start] == prefixXOR[end] {
				count += end - start - 1
			}
		}
	}
	return count
}
