import (
	"math"
	"sort"
)

const mod = 1e9 + 7

func numFactoredBinaryTrees(arr []int) int {
	sort.Ints(arr)
	s := make(map[int]struct{})
	for _, v := range arr {
		s[v] = struct{}{}
	}

	dp := make(map[int]int)
	for _, v := range arr {
		dp[v] = 1
	}

	for _, i := range arr {
		for _, j := range arr {
			if j > int(math.Sqrt(float64(i))) {
				break
			}

			_, ok := s[i/j]
			if i%j == 0 && ok {
				temp := (dp[j] * dp[i/j]) % mod
				if i/j == j {
					dp[i] = (dp[i] + temp) % mod
				} else {
					dp[i] = (dp[i] + temp*2) % mod
				}
			}
		}
	}

	res := 0
	for _, v := range dp {
		res = (res + v) % mod
	}

	return res
}
