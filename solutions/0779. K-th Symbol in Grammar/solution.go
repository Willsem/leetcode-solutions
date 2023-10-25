import "math"

func kthGrammar(n int, k int) int {
	if n == 1 {
		return 0
	}

	total := int(math.Pow(2, float64(n-1)))
	half := total / 2

	if k > half {
		return 1 - kthGrammar(n, k-half)
	}

	return kthGrammar(n-1, k)
}
