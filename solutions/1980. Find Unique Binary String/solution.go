func findDifferentBinaryString(nums []string) string {
	ans := ""
	for i := 0; i < len(nums); i++ {
		curr := nums[i][i]
		if curr == '0' {
			ans += "1"
		} else {
			ans += "0"
		}
	}

	return ans
}
