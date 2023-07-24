func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1.0 / myPow(x, -n)
	}

	result := 1.0
	for n != 0 {
		if n%2 == 1 {
			result *= x
			n--
		}

		x *= x
		n /= 2
	}
	return result
}
