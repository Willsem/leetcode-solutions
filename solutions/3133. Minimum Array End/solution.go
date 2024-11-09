func minEnd(n int, x int) int64 {
	res := x
	n--

	for mask := 1; n > 0; mask <<= 1 {
		if mask&x == 0 {
			res |= (n & 1) * mask
			n >>= 1
		}
	}

	return int64(res)
}
