func findMatrix(nums []int) [][]int {
	freq := make([]int, len(nums)+1)

	ans := make([][]int, 0)
	for _, num := range nums {
		if freq[num] >= len(ans) {
			ans = append(ans, make([]int, 0))
		}

		ans[freq[num]] = append(ans[freq[num]], num)
		freq[num]++
	}

	return ans
}
