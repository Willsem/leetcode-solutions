const MOD = 1e+9 + 7

func countOrders(n int) int {
	count := 1

	for i := 2; i <= n; i++ {
		count = (count * (2*i - 1) * i) % MOD
	}

	return count
}
