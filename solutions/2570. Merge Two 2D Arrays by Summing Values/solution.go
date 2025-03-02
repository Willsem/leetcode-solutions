import "math"

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	i1, i2 := 0, 0
	nums := make([][]int, 0)
	for i1 < len(nums1) || i2 < len(nums2) {
		num1, num2 := getFromNums(nums1, i1), getFromNums(nums2, i2)

		switch {
		case num1[0] == num2[0]:
			nums = append(nums, []int{num1[0], num1[1] + num2[1]})
			i1++
			i2++

		case num1[0] < num2[0]:
			nums = append(nums, num1)
			i1++

		case num2[0] < num1[0]:
			nums = append(nums, num2)
			i2++
		}
	}

	return nums
}

func getFromNums(nums [][]int, i int) []int {
	if i < len(nums) {
		return nums[i]
	}
	return []int{math.MaxInt64, 0}
}
