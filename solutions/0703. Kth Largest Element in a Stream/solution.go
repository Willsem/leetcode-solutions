import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
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

type KthLargest struct {
	heap *IntHeap
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	intHeap := &IntHeap{}
	heap.Init(intHeap)
	l := KthLargest{intHeap, k}

	for _, num := range nums {
		l.Add(num)
	}

	return l
}

func (l *KthLargest) Add(val int) int {
	if l.heap.Len() < l.k {
		heap.Push(l.heap, val)
	} else if val > (*l.heap)[0] {
		heap.Pop(l.heap)
		heap.Push(l.heap, val)
	}
	return (*l.heap)[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
