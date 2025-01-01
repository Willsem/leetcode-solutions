func maxScore(s string) int {
	countOnes, countZeros := 0, 0
	for _, c := range s {
		if c == '1' {
			countOnes++
		}
	}

	res := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '0' {
			countZeros++
		} else {
			countOnes--
		}

		res = max(res, countOnes+countZeros)
	}

	return res
}
