import "container/heap"

type HeapElement struct {
	avg   float64
	index int
}

type Heap []*HeapElement

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].avg > h[j].avg }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

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

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	h := &Heap{}
	for i, c := range classes {
		heap.Push(h, &HeapElement{
			avg:   avg(c),
			index: i,
		})
	}

	for i := 0; i < extraStudents; i++ {
		c := heap.Pop(h).(*HeapElement)
		classes[c.index][0]++
		classes[c.index][1]++
		c.avg = avg(classes[c.index])
		heap.Push(h, c)
	}

	sum := float64(0)
	for _, c := range classes {
		sum += float64(c[0]) / float64(c[1])
	}
	return sum / float64(len(classes))
}

func avg(c []int) float64 {
	return float64(c[0]+1)/float64(c[1]+1) - float64(c[0])/float64(c[1])
}
