func decrypt(code []int, k int) []int {
	res := make([]int, len(code))
	if k == 0 {
		return res
	}

	l, r, sum := 1, k, 0
	if k < 0 {
		l = len(code) + k
		r = len(code) - 1
	}

	for i := l; i <= r; i++ {
		sum += code[i]
	}

	for i := 0; i < len(code); i++ {
		res[i] = sum

		sum -= code[l]
		l++
		l %= len(code)
		r++
		r %= len(code)
		sum += code[r]
	}

	return res
}
