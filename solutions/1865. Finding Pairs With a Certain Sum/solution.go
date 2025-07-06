type FindSumPairs struct {
	nums1      []int
	nums2      []int
	nums2Count map[int]int
}

func Constructor(nums1 []int, nums2 []int) FindSumPairs {
	nums2Count := make(map[int]int)
	for _, num := range nums2 {
		nums2Count[num]++
	}

	return FindSumPairs{
		nums1:      nums1,
		nums2:      nums2,
		nums2Count: nums2Count,
	}
}

func (f *FindSumPairs) Add(i int, val int) {
	oldValue := f.nums2[i]
	f.nums2[i] += val
	f.nums2Count[f.nums2[i]]++

	f.nums2Count[oldValue]--
	if f.nums2Count[oldValue] == 0 {
		delete(f.nums2Count, oldValue)
	}
}

func (f *FindSumPairs) Count(tot int) int {
	count := 0
	for _, num := range f.nums1 {
		need := tot - num
		count += f.nums2Count[need]
	}
	return count
}

/**
 * Your FindSumPairs object will be instantiated and called as such:
 * obj := Constructor(nums1, nums2);
 * obj.Add(index,val);
 * param_2 := obj.Count(tot);
 */
