func findNumbers(nums []int) int {
	count := 0
	for _, num := range nums {
		if countDigits(num)%2 == 0 {
			count++
		}
	}
	return count
}

func countDigits(num int) int {
	count := 0
	for num > 0 {
		count++
		num /= 10
	}
	return count
}
