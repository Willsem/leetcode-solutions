func countLargestGroup(n int) int {
	sumToCount := make(map[int]int)

	for num := 1; num <= n; num++ {
		sumToCount[sumDigits(num)]++
	}

	maxCount, countOfMax := 0, 0
	for _, count := range sumToCount {
		if count > maxCount {
			maxCount = count
			countOfMax = 1
		} else if count == maxCount {
			countOfMax++
		}
	}

	return countOfMax
}

func sumDigits(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}
