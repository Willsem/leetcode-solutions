func longestCommonPrefix(arr1 []int, arr2 []int) int {
	prefix1 := make(map[int]int)
	for _, num := range arr1 {
		digits := parseInt(num)
		for i := 1; i <= len(digits); i++ {
			prefix1[makeKey(digits[:i])] = i
		}
	}

	maxLen := 0
	for _, num := range arr2 {
		digits := parseInt(num)
		for i := 1; i <= len(digits); i++ {
			maxLen = max(maxLen, prefix1[makeKey(digits[:i])])
		}
	}
	return maxLen
}

func parseInt(x int) []byte {
	digits := make([]byte, 0)
	for x > 0 {
		digits = append(digits, byte(x%10))
		x /= 10
	}
	for i := 0; i < len(digits)/2; i++ {
		digits[i], digits[len(digits)-i-1] = digits[len(digits)-i-1], digits[i]
	}
	return digits
}

func makeKey(prefix []byte) int {
	key := 0
	for _, d := range prefix {
		key = key*10 + int(d)
	}
	return key
}
