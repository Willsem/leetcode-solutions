type FenwickTree struct {
	tree []int
}

func NewFenwickTree(size int) *FenwickTree {
	return &FenwickTree{
		tree: make([]int, size+1),
	}
}

func (t *FenwickTree) update(index int, delta int) {
	for index < len(t.tree) {
		t.tree[index] += delta
		index += index & -index
	}
}

func (t *FenwickTree) query(index int) int {
	res := 0
	for index > 0 {
		res += t.tree[index]
		index -= index & -index
	}
	return res
}

func goodTriplets(nums1 []int, nums2 []int) int64 {
	n := len(nums1)
	pos2 := make([]int, n)
	for i, num := range nums2 {
		pos2[num] = i + 1
	}

	a := make([]int, n)
	for i, num := range nums1 {
		a[i] = pos2[num]
	}

	left := make([]int, n)
	right := make([]int, n)

	ft := NewFenwickTree(n)
	for i := 0; i < n; i++ {
		left[i] = ft.query(a[i] - 1)
		ft.update(a[i], 1)
	}

	ft = NewFenwickTree(n)
	for i := n - 1; i >= 0; i-- {
		right[i] = ft.query(n) - ft.query(a[i])
		ft.update(a[i], 1)
	}

	var res int64 = 0
	for i := 0; i < n; i++ {
		res += int64(left[i]) * int64(right[i])
	}
	return res
}
