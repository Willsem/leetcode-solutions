func validPartition(nums []int) bool {
	n := len(nums)

	dp := make([]bool, 3)
	dp[0] = true

	for i := 0; i < n; i++ {
		dpIndex := i + 1
		ans := false

		if i > 0 && nums[i] == nums[i-1] {
			ans = ans || dp[(dpIndex-2)%3]
		}
		if i > 1 && tripletEquals(nums[i], nums[i-1], nums[i-2]) {
			ans = ans || dp[(dpIndex-3)%3]
		}
		if i > 1 && tripletEquals(nums[i], nums[i-1]+1, nums[i-2]+2) {
			ans = ans || dp[(dpIndex-3)%3]
		}

		dp[dpIndex%3] = ans
	}

	return dp[n%3]
}

func tripletEquals(a, b, c int) bool {
	return a == b && a == c
}
