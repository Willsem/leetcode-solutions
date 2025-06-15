func maxDiff(num int) int {
	digits := make([]int, 0)
	for num > 0 {
		digits = append(digits, num%10)
		num /= 10
	}

	maxVal, minVal := 0, 0
	maxKey, minKey := digits[len(digits)-1], digits[len(digits)-1]
	maxReplace, minReplace := 9, 1

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 1 && digits[i] != 0 {
			minKey = digits[i]
			if i < len(digits)-1 {
				minReplace = 0
			}
			break
		}
	}

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			maxKey = digits[i]
			break
		}
	}

	for i := len(digits) - 1; i >= 0; i-- {
		maxDigit, minDigit := digits[i], digits[i]

		if digits[i] == maxKey {
			maxDigit = maxReplace
		}

		if digits[i] == minKey {
			minDigit = minReplace
		}

		maxVal = maxVal*10 + maxDigit
		minVal = minVal*10 + minDigit
	}

	return maxVal - minVal
}
