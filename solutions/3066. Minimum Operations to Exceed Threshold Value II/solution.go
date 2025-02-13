import "container/heap"

type Heap []int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] < h[j] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func NewHeap(data []int) *Heap {
	h := Heap(data)
	heap.Init(&h)
	return &h
}

func minOperations(nums []int, k int) int {
	h := Heap(nums)
	heap.Init(&h)

	countOperations := 0
	for nums[0] < k {
		x := heap.Pop(&h).(int)
		y := heap.Pop(&h).(int)
		heap.Push(&h, x*2+y)
		countOperations++
	}

	return countOperations
}
