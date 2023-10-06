func majorityElement(nums []int) []int {
	count1, count2 := 0, 0
	candidate1, candidate2 := 0, 0

	for _, num := range nums {
		if count1 == 0 && num != candidate2 {
			count1 = 1
			candidate1 = num
		} else if count2 == 0 && num != candidate1 {
			count2 = 1
			candidate2 = num
		} else if candidate1 == num {
			count1++
		} else if candidate2 == num {
			count2++
		} else {
			count1--
			count2--
		}
	}

	count1, count2 = 0, 0
	for _, num := range nums {
		if candidate1 == num {
			count1++
		} else if candidate2 == num {
			count2++
		}
	}

	result := make([]int, 0)
	threshold := len(nums) / 3
	if count1 > threshold {
		result = append(result, candidate1)
	}
	if count2 > threshold {
		result = append(result, candidate2)
	}

	return result
}
