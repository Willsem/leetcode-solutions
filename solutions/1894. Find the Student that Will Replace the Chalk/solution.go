func chalkReplacer(chalk []int, k int) int {
	chalkSum := 0
	for _, c := range chalk {
		chalkSum += c
		if chalkSum > k {
			break
		}
	}

	k %= chalkSum

	for i, c := range chalk {
		if k < c {
			return i
		}
		k -= c
	}

	return 0
}
