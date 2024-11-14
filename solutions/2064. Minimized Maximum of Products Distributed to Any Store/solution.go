func minimizedMaximum(n int, quantities []int) int {
	canDistribute := func(x int) bool {
		stores := 0
		for _, q := range quantities {
			stores += (q + x - 1) / x
		}
		return stores <= n
	}

	l, r := 1, 0
	for _, q := range quantities {
		if q > r {
			r = q
		}
	}

	result := 0
	for l <= r {
		x := (l + r) / 2
		if canDistribute(x) {
			result = x
			r = x - 1
		} else {
			l = x + 1
		}
	}

	return result
}
