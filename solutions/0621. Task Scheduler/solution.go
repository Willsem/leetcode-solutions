func leastInterval(tasks []byte, n int) int {
	counts := make(map[byte]int)
	maximum, maxCount := 0, 0

	for _, task := range tasks {
		counts[task]++
		if maximum == counts[task] {
			maxCount++
		} else if maximum < counts[task] {
			maximum = counts[task]
			maxCount = 1
		}
	}

	partCount := maximum - 1
	partLength := n - (maxCount - 1)
	emptySlots := partCount * partLength
	availableTasks := len(tasks) - maximum*maxCount
	idles := max(0, emptySlots-availableTasks)

	return len(tasks) + idles
}
