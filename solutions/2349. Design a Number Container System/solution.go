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

type NumberContainers struct {
	indexToNumber map[int]int
	numberToIndex map[int]*Heap
}

func Constructor() NumberContainers {
	return NumberContainers{
		indexToNumber: make(map[int]int),
		numberToIndex: make(map[int]*Heap),
	}
}

func (c *NumberContainers) Change(index int, number int) {
	if _, ok := c.numberToIndex[number]; !ok {
		c.numberToIndex[number] = &Heap{}
		heap.Init(c.numberToIndex[number])
	}

	heap.Push(c.numberToIndex[number], index)
	c.indexToNumber[index] = number
}

func (c *NumberContainers) Find(number int) int {
	if _, ok := c.numberToIndex[number]; !ok {
		return -1
	}

	for c.numberToIndex[number].Len() > 0 {
		index := heap.Pop(c.numberToIndex[number]).(int)
		if c.indexToNumber[index] == number {
			heap.Push(c.numberToIndex[number], index)
			return index
		}
	}

	return -1
}

/**
 * Your NumberContainers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Change(index,number);
 * param_2 := obj.Find(number);
 */
