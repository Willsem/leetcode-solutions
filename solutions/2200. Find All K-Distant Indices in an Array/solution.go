func findKDistantIndices(nums []int, key int, k int) []int {
	result := make([]int, 0)

	for i, num := range nums {
		if num == key {
			j := max(0, i-k)
			if len(result) > 0 {
				j = max(result[len(result)-1]+1, i-k)
			}

			for ; j <= min(len(nums)-1, i+k); j++ {
				result = append(result, j)
			}
		}
	}

	return result
}
