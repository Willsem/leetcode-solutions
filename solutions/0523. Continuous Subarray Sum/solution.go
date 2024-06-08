func checkSubarraySum(nums []int, k int) bool {
	prefixMod := 0
	mods := make(map[int]int)
	mods[0] = -1

	for i, num := range nums {
		prefixMod = (prefixMod + num) % k

		if _, ok := mods[prefixMod]; ok {
			if i-mods[prefixMod] > 1 {
				return true
			}
		} else {
			mods[prefixMod] = i
		}
	}

	return false
}
