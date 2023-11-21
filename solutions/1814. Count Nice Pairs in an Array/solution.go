const mod = 1e9 + 7

func countNicePairs(nums []int) int {
	for i := range nums {
		nums[i] -= rev(nums[i])
	}

	prev := make(map[int]int)
	count := 0
	for _, num := range nums {
		if v, ok := prev[num]; ok {
			count += v
			count %= mod
		}

		prev[num]++
	}

	return count
}

func rev(num int) int {
	res := 0
	for num > 0 {
		res *= 10
		res += num % 10
		num /= 10
	}
	return res
}
