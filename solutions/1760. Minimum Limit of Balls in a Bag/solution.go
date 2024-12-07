func canDivide(nums []int, maxBalls int, maxOperations int) bool {
	ops := 0
	for _, n := range nums {
		ops += (n+maxBalls-1)/maxBalls - 1
		if ops > maxOperations {
			return false
		}
	}
	return true
}

func minimumSize(nums []int, maxOperations int) int {
	l, r := 1, 0
	for _, n := range nums {
		if n > r {
			r = n
		}
	}
	ressult := r

	for l < r {
		mid := (l + r) / 2
		if canDivide(nums, mid, maxOperations) {
			r = mid
			ressult = r
		} else {
			l = mid + 1
		}
	}

	return ressult
}
