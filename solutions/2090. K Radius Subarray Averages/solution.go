import "math"

func getAverages(nums []int, k int) []int {
	result := make([]int, len(nums))
	for i := range result {
		result[i] = -1
	}

	if len(nums) < k*2+1 {
		return result
	}

	sum := 0
	count := k*2 + 1

	for i := 0; i < k*2; i++ {
		sum += nums[i]
	}

	for i := k; i < len(nums)-k; i++ {
		sum += nums[i+k]
		result[i] = int(math.Floor(float64(sum) / float64(count)))
		sum -= nums[i-k]
	}

	return result
}
