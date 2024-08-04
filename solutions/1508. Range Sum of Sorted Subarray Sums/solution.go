import "sort"

const mod = 1e9 + 7

func rangeSum(nums []int, n int, left int, right int) int {
	subarrays := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			subarrays = append(subarrays, sum)
		}
	}

	sort.Ints(subarrays)

	sum := 0
	for i := left - 1; i <= right-1; i++ {
		sum = (sum + subarrays[i]) % mod
	}

	return sum
}
