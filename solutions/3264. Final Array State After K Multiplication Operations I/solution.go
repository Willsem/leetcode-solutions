import "container/heap"

type HeapElement struct {
	value int
	index int
}

type Heap []*HeapElement

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	if h[i].value == h[j].value {
		return h[i].index < h[j].index
	}
	return h[i].value < h[j].value
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x any) {
	*h = append(*h, x.(*HeapElement))
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func getFinalState(nums []int, k int, multiplier int) []int {
	h := &Heap{}
	for i, num := range nums {
		*h = append(*h, &HeapElement{
			value: num,
			index: i,
		})
	}

	heap.Init(h)
	for ; k > 0; k-- {
		curr := heap.Pop(h).(*HeapElement)
		curr.value *= multiplier
		heap.Push(h, curr)
	}

	for h.Len() > 0 {
		curr := heap.Pop(h).(*HeapElement)
		nums[curr.index] = curr.value
	}

	return nums
}
