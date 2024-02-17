import "container/heap"

type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func furthestBuilding(heights []int, bricks int, ladders int) int {
	h := &MinHeap{}

	for i := 1; i < len(heights); i++ {
		diff := heights[i] - heights[i-1]
		if diff <= 0 {
			continue
		}

		heap.Push(h, diff)
		if h.Len() > ladders {
			bricks -= heap.Pop(h).(int)
		}

		if bricks < 0 {
			return i - 1
		}
	}

	return len(heights) - 1
}
