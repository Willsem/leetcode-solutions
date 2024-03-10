func intersection(nums1 []int, nums2 []int) []int {
	set1 := make(map[int]struct{})
	for _, num := range nums1 {
		set1[num] = struct{}{}
	}

	intr := make(map[int]struct{})
	for _, num := range nums2 {
		if _, ok := set1[num]; ok {
			intr[num] = struct{}{}
		}
	}

	res := make([]int, 0, len(intr))
	for num := range intr {
		res = append(res, num)
	}

	return res
}
