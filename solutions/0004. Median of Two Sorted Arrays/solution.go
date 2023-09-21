func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l := len(nums1) + len(nums2)
	merged := make([]int, 0, l)
	endIndex := l / 2

	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) && len(merged) <= endIndex {
		if nums1[i] < nums2[j] {
			merged = append(merged, nums1[i])
			i++
		} else {
			merged = append(merged, nums2[j])
			j++
		}
	}

	for len(merged) <= endIndex {
		for i < len(nums1) {
			merged = append(merged, nums1[i])
			i++
		}
		for j < len(nums2) {
			merged = append(merged, nums2[j])
			j++
		}
	}

	if l%2 != 0 {
		return float64(merged[endIndex])
	}

	return (float64(merged[endIndex-1]) + float64(merged[endIndex])) / 2.0
}
