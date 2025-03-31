import "sort"

func putMarbles(weights []int, k int) int64 {
	pairWeights := make([]int, len(weights)-1)

	for i := 0; i < len(weights)-1; i++ {
		pairWeights[i] = weights[i] + weights[i+1]
	}

	sort.Ints(pairWeights)

	answer := 0
	for i := 0; i < k-1; i++ {
		answer += pairWeights[n-2-i] - pairWeights[i]
	}

	return int64(answer)
}
