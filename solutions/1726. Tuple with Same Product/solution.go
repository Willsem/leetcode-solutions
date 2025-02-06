func tupleSameProduct(nums []int) int {
	countPowers := make(map[int]int)

	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			countPowers[nums[i]*nums[j]]++
		}
	}

	result := 0
	for _, count := range countPowers {
		totalPairs := ((count - 1) * count) / 2
		result += 8 * totalPairs
	}
	return result
}
