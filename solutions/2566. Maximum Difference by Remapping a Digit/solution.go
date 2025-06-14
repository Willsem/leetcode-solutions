func minMaxDifference(num int) int {
	digits := make([]int, 0)
	for num > 0 {
		digits = append(digits, num%10)
		num /= 10
	}

	maxKey := 0
	for i := len(digits) - 1; i >= 0; i-- {
		maxKey = digits[i]
		if maxKey != 9 {
			break
		}
	}

	minKey := digits[len(digits)-1]
	maxVal, minVal := 0, 0

	for i := len(digits) - 1; i >= 0; i-- {
		maxDigit := digits[i]
		minDigit := digits[i]

		if digits[i] == maxKey {
			maxDigit = 9
		}

		if digits[i] == minKey {
			minDigit = 0
		}

		maxVal = maxVal*10 + maxDigit
		minVal = minVal*10 + minDigit
	}

	return maxVal - minVal
}
