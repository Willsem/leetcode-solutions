import "math/bits"

func canSortArray(nums []int) bool {
	n := len(nums)
	values := make([]int, len(nums))
	copy(values, nums)

	for i := 0; i < n-1; i++ {
		if values[i] <= values[i+1] {
			continue
		}

		if bits.OnesCount(uint(values[i])) == bits.OnesCount(uint(values[i+1])) {
			values[i], values[i+1] = values[i+1], values[i]
		} else {
			return false
		}
	}

	for i := n - 1; i > 0; i-- {
		if values[i] >= values[i-1] {
			continue
		}

		if bits.OnesCount(uint(values[i])) == bits.OnesCount(uint(values[i-1])) {
			values[i], values[i-1] = values[i-1], values[i]
		} else {
			return false
		}
	}

	return true
}
