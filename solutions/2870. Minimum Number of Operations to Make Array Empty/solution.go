import "math"

func minOperations(nums []int) int {
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}

	ans := 0
	for _, count := range counts {
		if count == 1 {
			return -1
		}

		ans += int(math.Ceil(float64(count) / 3.0))
	}

	return ans
}
