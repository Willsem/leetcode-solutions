const mod = 1e9 + 7

func countVowelPermutation(n int) int {
	a, e, i, o, u := 1, 1, 1, 1, 1

	for j := 1; j < n; j++ {
		a_next := e
		e_next := (a + i) % mod
		i_next := (a + e + o + u) % mod
		o_next := (i + u) % mod
		u_next := a

		a, e, i, o, u = a_next, e_next, i_next, o_next, u_next
	}

	return (a + e + i + o + u) % mod
}
