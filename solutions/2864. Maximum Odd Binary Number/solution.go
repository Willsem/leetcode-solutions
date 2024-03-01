func maximumOddBinaryNumber(s string) string {
	countOnes := 0
	for _, c := range s {
		if c == '1' {
			countOnes++
		}
	}

	res := make([]byte, len(s))
	for i := 0; i < countOnes-1; i++ {
		res[i] = '1'
	}
	for i := countOnes - 1; i < len(res)-1; i++ {
		res[i] = '0'
	}
	res[len(res)-1] = '1'

	return string(res)
}
