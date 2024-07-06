func passThePillow(n int, time int) int {
	fullPass := time / (n - 1)
	excess := time % (n - 1)
	if fullPass%2 == 0 {
		return excess + 1
	} else {
		return n - excess
	}
}
