func subarraysDivByK(nums []int, k int) int {
	prefixMod := 0
	count := 0
	mods := make(map[int]int)
	mods[0] = 1

	for _, num := range nums {
		prefixMod = (prefixMod + num%k + k) % k
		count += mods[prefixMod]
		mods[prefixMod]++
	}

	return count
}
