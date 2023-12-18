func maxProductDifference(nums []int) int {
	max, secondMax := 0, 0
	min, secondMin := 10001, 10001

	for _, num := range nums {
		if num > max {
			max, secondMax = num, max
		} else if num > secondMax {
			secondMax = num
		}

		if num < min {
			min, secondMin = num, min
		} else if num < secondMin {
			secondMin = num
		}
	}

	return max*secondMax - min*secondMin
}
