func minimumReplacement(nums []int) int64 {
	var operations int64 = 0
	var prev_bound int64 = int64(nums[len(nums)-1])

	for i := len(nums) - 2; i >= 0; i-- {
		var num int64 = int64(nums[i])
		var no_of_times int64 = (num + prev_bound - 1) / prev_bound
		operations += no_of_times - 1
		prev_bound = num / no_of_times
	}

	return operations
}
