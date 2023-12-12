func maxProduct(nums []int) int {
	max, secondMax := 0, 0
	for _, num := range nums {
		if num > max {
			max, secondMax = num, max
		} else if num > secondMax {
			secondMax = num
		}
	}

	return (max - 1) * (secondMax - 1)
}
