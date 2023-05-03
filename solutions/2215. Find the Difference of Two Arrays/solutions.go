func findDifference(nums1 []int, nums2 []int) [][]int {
	set1 := make(map[int]struct{}, len(nums1))
	set2 := make(map[int]struct{}, len(nums2))

	setAnswer1 := make(map[int]struct{}, 0)
	setAnswer2 := make(map[int]struct{}, 0)

	for _, v := range nums1 {
		set1[v] = struct{}{}
	}
	for _, v := range nums2 {
		set2[v] = struct{}{}
		if _, ok := set1[v]; !ok {
			setAnswer2[v] = struct{}{}
		}
	}
	for _, v := range nums1 {
		if _, ok := set2[v]; !ok {
			setAnswer1[v] = struct{}{}
		}
	}

	answer1 := make([]int, 0, len(setAnswer1))
	answer2 := make([]int, 0, len(setAnswer2))

	for key := range setAnswer1 {
		answer1 = append(answer1, key)
	}
	for key := range setAnswer2 {
		answer2 = append(answer2, key)
	}

	return [][]int{answer1, answer2}
}
