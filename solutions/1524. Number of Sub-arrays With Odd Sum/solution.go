const mod = 1_000_000_007

func numOfSubarrays(arr []int) int {
	count, prefixSum := 0, 0
	odds, evens := 0, 1

	for _, num := range arr {
		prefixSum += num
		if prefixSum%2 == 0 {
			count += odds
			evens++
		} else {
			count += evens
			odds++
		}

		count %= mod
	}

	return count
}
