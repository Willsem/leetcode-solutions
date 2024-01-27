const mod = 1e9 + 7

func kInversePairs(n, k int) int {
	prev, curr := make([]int, k+1), make([]int, k+1)
	prev[0], curr[0] = 1, 1

	for N := 1; N <= n; N++ {
		for K := 0; K <= k; K++ {
			temp := 0
			if K > 0 {
				temp = curr[K-1]
			}
			curr[K] = (prev[K] + temp) % mod

			temp = 0
			if K >= N {
				temp = prev[K-N]
			}
			curr[K] = (curr[K] + mod - temp) % mod
		}
		copy(prev, curr)
	}

	return curr[k]
}
