func sortColors(nums []int) {
	counts := []int{0, 0, 0}
	for _, num := range nums {
		counts[num]++
	}

	numsI := 0
	for value, count := range counts {
		for i := 0; i < count; i++ {
			nums[numsI] = value
			numsI++
		}
	}
}
