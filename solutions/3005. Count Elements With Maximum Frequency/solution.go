func maxFrequencyElements(nums []int) int {
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}

	max, countMaxs := 0, 0
	for _, count := range counts {
		if count == max {
			countMaxs++
		} else if count > max {
			max = count
			countMaxs = 1
		}
	}

	return countMaxs * max
}
