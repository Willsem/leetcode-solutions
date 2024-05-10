import "container/heap"

type fraction struct {
	a     int
	b     int
	index int
	value float64
}

type Heap []*fraction

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i].value < h[j].value
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *Heap) Push(x any) {
	n := len(*h)
	item := x.(*fraction)
	item.index = n
	*h = append(*h, item)
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*h = old[0 : n-1]
	return item
}

func (h *Heap) update(item *fraction, a int, b int, value float64) {
	item.a = a
	item.b = b
	item.value = value
	heap.Fix(h, item.index)
}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	h := make(Heap, len(arr)-1)

	for i := 0; i < len(arr)-1; i++ {
		h[i] = &fraction{
			a:     i,
			b:     len(arr) - 1,
			index: i,
			value: float64(arr[i]) / float64(arr[len(arr)-1]),
		}
	}

	heap.Init(&h)

	for k > 1 {
		if h[0].a < h[0].b-1 {
			h.update(h[0], h[0].a, h[0].b-1, float64(arr[h[0].a])/float64(arr[h[0].b-1]))
		} else {
			heap.Pop(&h)
		}
		k--
	}

	ans := heap.Pop(&h).(*fraction)
	return []int{arr[ans.a], arr[ans.b]}
}
