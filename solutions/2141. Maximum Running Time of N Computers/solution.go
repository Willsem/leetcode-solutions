func maxRunTime(n int, batteries []int) int64 {
	sum := 0
	for _, b := range batteries {
		sum += b
	}

	left, right := 1, sum/n

	for left < right {
		target := right - (right-left)/2

		extra := 0
		for _, power := range batteries {
			extra += min(power, target)
		}

		if extra/n >= target {
			left = target
		} else {
			right = target - 1
		}
	}

	return int64(left)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
