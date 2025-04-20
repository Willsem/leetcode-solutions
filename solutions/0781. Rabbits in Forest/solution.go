func numRabbits(answers []int) int {
	count := make(map[int]int)
	for _, ans := range answers {
		count[ans]++
	}

	result := 0
	for answer, count := range count {
		result += ((answer + count) / (answer + 1)) * (answer + 1)
	}
	return result
}
