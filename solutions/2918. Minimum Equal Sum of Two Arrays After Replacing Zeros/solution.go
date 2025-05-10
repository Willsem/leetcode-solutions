func minSum(nums1 []int, nums2 []int) int64 {
	sum1, zero1 := calcArray(nums1)
	sum2, zero2 := calcArray(nums2)

	if zero1 == 0 && sum2 > sum1 || zero2 == 0 && sum1 > sum2 {
		return -1
	}

	return int64(max(sum1, sum2))
}

func calcArray(nums []int) (int, int) {
	sum, zero := 0, 0
	for _, num := range nums {
		sum += num
		if num == 0 {
			zero++
			sum++
		}
	}
	return sum, zero
}
