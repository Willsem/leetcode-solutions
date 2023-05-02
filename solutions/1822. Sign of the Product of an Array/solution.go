func arraySign(nums []int) int {
	x := 1
	for _, v := range nums {
		if v == 0 {
			return 0
		} else if v < 0 {
			x *= -1
		}
	}

	return x
}
