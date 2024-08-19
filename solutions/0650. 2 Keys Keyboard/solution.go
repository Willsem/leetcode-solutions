func minSteps(n int) int {
	res := 0
	d := 2

	for n > 1 {
		for n%d == 0 {
			res += d
			n /= d
		}
		d++
	}

	return res
}
