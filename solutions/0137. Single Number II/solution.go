func singleNumber(nums []int) int {
	ones := 0
	twos := 0

	for _, number := range nums {
		ones ^= (number & ^twos)
		twos ^= (number & ^ones)
	}

	return ones
}
