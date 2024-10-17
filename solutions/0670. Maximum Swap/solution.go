func maximumSwap(num int) int {
	digits := parseNum(num)
	maxNum := num
	for i := 0; i < len(digits); i++ {
		for j := i + 1; j < len(digits); j++ {
			digits[i], digits[j] = digits[j], digits[i]
			maxNum = max(maxNum, collectNum(digits))
			digits[i], digits[j] = digits[j], digits[i]
		}
	}
	return maxNum
}

func parseNum(num int) []byte {
	parsed := make([]byte, 0)
	for num > 0 {
		parsed = append(parsed, byte(num%10))
		num /= 10
	}
	return parsed
}

func collectNum(digits []byte) int {
	collected := 0
	for i := len(digits) - 1; i >= 0; i-- {
		collected = collected*10 + int(digits[i])
	}
	return collected
}
