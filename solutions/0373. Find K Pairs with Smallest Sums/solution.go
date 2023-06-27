import "container/heap"

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	n := len(nums1)
	m := len(nums2)
	h := make(Heap, n)

	for i := 0; i < n; i++ {
		h[i] = [3]int{nums1[i], nums2[0], 0}
	}

	heap.Init(&h)

	res := make([][]int, min(m*n, k))
	for i := 0; i < min(m*n, k); i++ {
		pair := heap.Pop(&h).([3]int)
		res[i] = []int{pair[0], pair[1]}
		if pair[2] != m-1 {
			heap.Push(&h, [3]int{pair[0], nums2[pair[2]+1], pair[2] + 1})
		}
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// Heap element format: [num1, num2, idx2]. Sorted based on sum of num1 + num2

type Heap [][3]int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h Heap) Less(i, j int) bool { return h[i][0]+h[i][1] < h[j][0]+h[j][1] }

func (h *Heap) Push(x any) {
	*h = append(*h, x.([3]int))
}
func (h *Heap) Pop() any {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}
