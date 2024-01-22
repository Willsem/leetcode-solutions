import "math"

func findErrorNums(nums []int) []int {
	ans := []int{}

	for i := range nums {
		num := int(math.Abs(float64(nums[i])))
		if nums[num-1] < 0 {
			ans = append(ans, num)
		} else {
			nums[num-1] *= -1
		}
	}

	for i := range nums {
		if nums[i] > 0 {
			ans = append(ans, i+1)
			break
		}
	}

	return ans
}
