func maximumSum(nums []int) int {
	sumOfDigitsToNums := make(map[int][]int)

	for _, num := range nums {
		sum := sumOfDigits(num)
		sumOfDigitsToNums[sum] = append(sumOfDigitsToNums[sum], num)
	}

	maxSum := -1
	for _, nums := range sumOfDigitsToNums {
		if len(nums) < 2 {
			continue
		}

		firstMax, secondMax := nums[0], nums[1]
		if secondMax > firstMax {
			firstMax, secondMax = secondMax, firstMax
		}

		for _, num := range nums[2:] {
			if num > firstMax {
				firstMax, secondMax = num, firstMax
			} else if num > secondMax {
				secondMax = num
			}
		}

		maxSum = max(maxSum, firstMax+secondMax)
	}

	return maxSum
}

func sumOfDigits(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}
