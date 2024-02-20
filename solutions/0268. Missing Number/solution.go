func missingNumber(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	n := len(nums)
	return (n*(n+1))/2 - sum
}
