const mod = 1e9 + 7

func i(char rune) int {
	return int(char - 'a')
}

func lengthAfterTransformations(s string, t int) int {
	counts := make([]int, 26)
	for _, c := range s {
		counts[i(c)]++
	}

	for range t {
		next := make([]int, 26)
		next[i('a')] = counts[i('z')]
		next[i('b')] = (counts[i('z')] + counts[i('a')]) % mod
		for c := 'c'; c <= 'z'; c++ {
			next[i(c)] = counts[i(c)-1]
		}
		counts = next
	}

	sum := 0
	for _, count := range counts {
		sum = (sum + count) % mod
	}
	return sum
}
