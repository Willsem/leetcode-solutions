import "sort"

func putMarbles(weights []int, k int) int64 {
	n := len(weights)
	pairWeights := make([]int, n-1)

	for i := 0; i < n-1; i++ {
		pairWeights[i] = weights[i] + weights[i+1]
	}

	sort.Ints(pairWeights)

	answer := 0
	for i := 0; i < k-1; i++ {
		answer += pairWeights[n-2-i] - pairWeights[i]
	}

	return int64(answer)
}
