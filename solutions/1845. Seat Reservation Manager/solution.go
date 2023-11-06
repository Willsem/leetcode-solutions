import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type SeatManager struct {
	heap *IntHeap
}

func Constructor(n int) SeatManager {
	h := &IntHeap{}
	heap.Init(h)

	for i := 1; i <= n; i++ {
		heap.Push(h, i)
	}

	return SeatManager{
		heap: h,
	}
}

func (m *SeatManager) Reserve() int {
	return heap.Pop(m.heap).(int)
}

func (m *SeatManager) Unreserve(seatNumber int) {
	heap.Push(m.heap, seatNumber)
}

/**
 * Your SeatManager object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Reserve();
 * obj.Unreserve(seatNumber);
 */
