func findLucky(arr []int) int {
	counts := make(map[int]int)
	for _, num := range arr {
		counts[num]++
	}

	maxLuckyNum := -1
	for num, count := range counts {
		if num == count && num > maxLuckyNum {
			maxLuckyNum = num
		}
	}
	return maxLuckyNum
}
