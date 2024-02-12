func majorityElement(nums []int) int {
	curr, count := 0, 0

	for _, num := range nums {
		if count == 0 {
			curr = num
			count++
		} else if curr == num {
			count++
		} else {
			count--
		}
	}

	return curr
}
