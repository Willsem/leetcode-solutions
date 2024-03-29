func countSubarrays(nums []int, k int) int64 {
	max := 0
	for _, num := range nums {
		if max < num {
			max = num
		}
	}

	left, count, ans := 0, 0, 0
	for right := range nums {
		if nums[right] == max {
			count++
		}

		for count == k {
			if nums[left] == max {
				count--
			}
			left++
		}

		ans += left
	}

	return int64(ans)
}
