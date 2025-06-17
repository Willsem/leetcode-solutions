const (
	mx  = 1e5
	mod = 1e9 + 7
)

func pow(x, n int) int {
	res := 1
	for n > 0 {
		if n&1 == 1 {
			res = res * x % mod
		}
		x = x * x % mod
		n >>= 1
	}
	return res
}

func countGoodArrays(n int, m int, k int) int {
	fact := make([]int, mx)
	invFact := make([]int, mx)

	fact[0] = 1
	for i := 1; i < mx; i++ {
		fact[i] = fact[i-1] * i % mod
	}

	invFact[mx-1] = pow(fact[mx-1], mod-2)
	for i := int(mx - 1); i > 0; i-- {
		invFact[i-1] = invFact[i] * i % mod
	}

	comb := func(n, m int) int {
		return fact[n] * invFact[m] % mod * invFact[n-m] % mod
	}

	return comb(n-1, k) * m % mod * pow(m-1, n-k-1) % mod
}
