func minBitFlips(start int, goal int) int {
	count := 0
	for start > 0 || goal > 0 {
		startBit, goalBit := start%2, goal%2
		if startBit != goalBit {
			count++
		}
		start, goal = start/2, goal/2
	}
	return count
}
