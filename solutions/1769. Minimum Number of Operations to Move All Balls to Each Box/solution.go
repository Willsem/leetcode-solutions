import "strconv"

func minOperations(boxes string) []int {
	result := make([]int, len(boxes))

	ballsLeft, movesLeft, ballsRight, movesRight := 0, 0, 0, 0
	for i := range boxes {
		result[i] += movesLeft

		balls, _ := strconv.Atoi(string(boxes[i]))
		ballsLeft += balls
		movesLeft += ballsLeft

		j := len(boxes) - i - 1
		result[j] += movesRight

		balls, _ = strconv.Atoi(string(boxes[j]))
		ballsRight += balls
		movesRight += ballsRight
	}

	return result
}
