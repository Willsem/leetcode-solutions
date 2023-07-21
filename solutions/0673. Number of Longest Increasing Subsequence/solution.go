func findNumberOfLIS(nums []int) int {
	n := len(nums)

	length := make([]int, n)
	count := make([]int, n)

	var calculate func(i int)
	calculate = func(i int) {
		if length[i] != 0 {
			return
		}

		length[i] = 1
		count[i] = 1

		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				calculate(j)

				if length[j]+1 > length[i] {
					length[i] = length[j] + 1
					count[i] = 0
				}

				if length[j]+1 == length[i] {
					count[i] += count[j]
				}
			}
		}
	}

	maxLength := 0
	result := 0
	for i := range nums {
		calculate(i)
		if length[i] > maxLength {
			maxLength = length[i]
		}
	}

	for i := range nums {
		if length[i] == maxLength {
			result += count[i]
		}
	}

	return result
}
