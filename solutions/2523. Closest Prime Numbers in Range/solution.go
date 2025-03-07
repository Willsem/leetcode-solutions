func closestPrimes(left int, right int) []int {
	num1, num2, prev := -1, -1, -1

	for i := left; i <= right; i++ {
		if isPrime(i) {
			if num2 == -1 {
				num2 = i
			} else if num1 == -1 {
				num1, num2 = num2, i
			} else {
				if num2-num1 > i-num2 {
					num1, num2 = num2, i
				}
				if num2-num1 > i-prev {
					num1, num2 = prev, i
				}
			}
			prev = i
		}
	}

	if num1 == -1 || num2 == -1 {
		return []int{-1, -1}
	}

	return []int{num1, num2}
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}

	if num == 2 || num == 3 {
		return true
	}

	if num%2 == 0 {
		return false
	}

	for div := 3; div*div <= num; div += 2 {
		if num%div == 0 {
			return false
		}
	}

	return true
}
