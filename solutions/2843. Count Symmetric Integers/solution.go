func countSymmetricIntegers(low int, high int) int {
	count := 0
	for num := low; num <= high; num++ {
		if isSymmetric(num) {
			count++
		}
	}
	return count
}

func isSymmetric(num int) bool {
	digits := make([]int, 0)
	for num > 0 {
		digits = append(digits, num%10)
		num /= 10
	}

	if len(digits)%2 != 0 {
		return false
	}

	sum1, sum2 := 0, 0
	for i, digit := range digits {
		if i < len(digits)/2 {
			sum1 += digit
		} else {
			sum2 += digit
		}
	}

	return sum1 == sum2
}
