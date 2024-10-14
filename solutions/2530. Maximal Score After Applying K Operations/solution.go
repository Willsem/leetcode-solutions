import (
	"container/heap"
	"math"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] } // Max-heap property
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxKelements(nums []int, k int) int64 {
	intHeap := &IntHeap{}
	heap.Init(intHeap)

	for _, num := range nums {
		heap.Push(intHeap, num)
	}

	result := 0
	for i := 0; i < k; i++ {
		num := heap.Pop(intHeap).(int)
		result += num
		heap.Push(intHeap, int(math.Ceil(float64(num)/3.0)))
	}

	return int64(result)
}
