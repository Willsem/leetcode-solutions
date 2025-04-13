const mod int64 = 1e9 + 7

func countGoodNumbers(n int64) int {
	return int(power(5, (n+1)/2) * power(4, n/2) % mod)
}

func power(x, y int64) int64 {
	result := int64(1)
	mult := x % mod
	for y > 0 {
		if y%2 == 1 {
			result = (result * mult) % mod
		}
		mult = (mult * mult) % mod
		y /= 2
	}
	return result
}
