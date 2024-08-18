func nthUglyNumber(n int) int {
	nums := make([]int, n)
	nums[0] = 1

	index2, index3, index5 := 0, 0, 0
	next2, next3, next5 := 2, 3, 5

	for i := 1; i < n; i++ {
		nextNum := min(min(next2, next3), next5)
		nums[i] = nextNum

		if nextNum == next2 {
			index2++
			next2 = nums[index2] * 2
		}
		if nextNum == next3 {
			index3++
			next3 = nums[index3] * 3
		}
		if nextNum == next5 {
			index5++
			next5 = nums[index5] * 5
		}
	}

	return nums[n-1]
}
