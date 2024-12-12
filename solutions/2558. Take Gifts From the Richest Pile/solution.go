import (
	"container/heap"
	"math"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
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

func NewIntHeap(data []int) *IntHeap {
	h := IntHeap(data)
	heap.Init(&h)
	return &h
}

func pickGifts(gifts []int, k int) int64 {
	h := NewIntHeap(gifts)

	for ; k > 0; k-- {
		gift := heap.Pop(h).(int)
		heap.Push(h, int(math.Floor(math.Sqrt(float64(gift)))))
	}

	count := 0
	for _, gift := range gifts {
		count += gift
	}
	return int64(count)
}
