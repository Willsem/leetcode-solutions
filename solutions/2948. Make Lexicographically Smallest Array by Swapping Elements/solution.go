import "slices"

func lexicographicallySmallestArray(nums []int, limit int) []int {
	n := len(nums)
	indices := make([]int, n)
	for i := range indices {
		indices[i] = i
	}

	slices.SortFunc(indices, func(i, j int) int {
		return nums[i] - nums[j]
	})

	result := make([]int, n)
	for i := 0; i < n; {
		j := i + 1

		for j < n && nums[indices[j]]-nums[indices[j-1]] <= limit {
			j++
		}

		subset := slices.Clone(indices[i:j])
		slices.Sort(subset)

		for k := i; k < j; k++ {
			result[subset[k-i]] = nums[indices[k]]
		}

		i = j
	}

	return result
}
