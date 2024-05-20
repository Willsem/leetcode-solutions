func subsetXORSum(nums []int) int {
	res := 0

	var subsets func(curr []int, i int)
	subsets = func(curr []int, i int) {
		if i == len(nums) {
			xor := 0
			for _, num := range curr {
				xor ^= num
			}
			res += xor
			return
		}

		subsets(curr, i+1)
		curr = append(curr, nums[i])
		subsets(curr, i+1)
		curr = curr[:len(curr)-1]
	}

	subsets([]int{}, 0)
	return res
}
