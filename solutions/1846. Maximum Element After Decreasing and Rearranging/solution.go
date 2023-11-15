import "math"

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	counts := make([]int, len(arr)+1)

	for _, num := range arr {
		counts[int(math.Min(float64(num), float64(len(arr))))]++
	}

	ans := 1
	for num := 2; num <= len(arr); num++ {
		ans = int(math.Min(float64(ans+counts[num]), float64(num)))
	}

	return ans
}
