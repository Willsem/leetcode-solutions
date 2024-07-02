func intersect(nums1 []int, nums2 []int) []int {
	countsNums2 := make(map[int]int)
	for _, num := range nums2 {
		countsNums2[num]++
	}

	res := make([]int, 0)
	for _, num := range nums1 {
		if _, ok := countsNums2[num]; ok {
			countsNums2[num]--
			if countsNums2[num] == 0 {
				delete(countsNums2, num)
			}
			res = append(res, num)
		}
	}
	return res
}
