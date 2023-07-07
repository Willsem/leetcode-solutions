func maxConsecutiveAnswers(answerKey string, k int) int {
	return max(maxConsecutiveAnswersBySymbol(answerKey, k, 'F'), maxConsecutiveAnswersBySymbol(answerKey, k, 'T'))
}

func maxConsecutiveAnswersBySymbol(answerKey string, k int, symbol byte) int {
	left := 0
	changes := 0
	maxLen := 0
	currLen := 0

	for right := range answerKey {
		if answerKey[right] != symbol {
			for changes >= k {
				if answerKey[left] != symbol {
					changes--
				}

				left++
				currLen--
			}

			changes++
		}

		currLen++
		maxLen = max(maxLen, currLen)
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
