import "sort"

func maximumHappinessSum(happiness []int, k int) int64 {
	sort.Slice(happiness, func(i, j int) bool {
		return happiness[i] > happiness[j]
	})

	result := 0
	for i := 0; i < k; i++ {
		result += max(happiness[i]-i, 0)
	}

	return int64(result)
}
