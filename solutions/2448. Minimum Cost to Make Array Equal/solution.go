func minCost(nums []int, cost []int) int64 {
	getCost := func(base int) int {
		sum := 0
		for i := range nums {
			sum += cost[i] * abs(nums[i]-base)
		}
		return sum
	}

	left := nums[0]
	right := nums[0]
	for _, v := range nums {
		if left > v {
			left = v
		}

		if right < v {
			right = v
		}
	}

	answer := getCost(nums[0])
	for left < right {
		mid := (left + right) / 2
		costCurr := getCost(mid)
		costNext := getCost(mid + 1)
		answer = min(costCurr, costNext)

		if costCurr > costNext {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return int64(answer)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
