const mod = 1e9 + 7

func knightDialer(n int) int {
	if n == 1 {
		return 10
	}

	a, b, c, d := 4, 2, 2, 1

	for i := 0; i < n-1; i++ {
		ta, tb, tc, td := a, b, c, d
		a = ((2*tb)%mod + (2*tc)%mod) % mod
		b = ta
		c = (ta + (2*td)%mod) % mod
		d = tc
	}

	return (((a+b)%mod+c)%mod + d) % mod
}
