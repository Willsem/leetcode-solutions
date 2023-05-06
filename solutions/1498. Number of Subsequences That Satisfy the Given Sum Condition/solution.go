import "sort"

const mod = 1_000_000_007

func numSubseq(nums []int, target int) int {
	sort.Ints(nums)

	powers := make([]int, len(nums))
	powers[0] = 1
	for i := 1; i < len(powers); i++ {
		powers[i] = (2 * powers[i-1]) % mod
	}

	result := 0
	l, r := 0, len(nums)-1
	for l <= r {
		if nums[l]+nums[r] > target {
			r--
		} else {
			result += powers[r-l]
			result %= mod
			l++
		}
	}

	return result
}
