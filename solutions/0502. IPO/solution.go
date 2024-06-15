import (
	"container/heap"
	"sort"
)

type project struct {
	capital int
	profit  int
}

type projectHeap []int

func (h projectHeap) Len() int            { return len(h) }
func (h projectHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h projectHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *projectHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *projectHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	projects := make([]project, len(profits))
	for i := 0; i < len(profits); i++ {
		projects[i] = project{capital[i], profits[i]}
	}

	sort.Slice(projects, func(i, j int) bool {
		return projects[i].capital < projects[j].capital
	})

	maxHeap := &projectHeap{}
	heap.Init(maxHeap)

	for i := 0; i < len(projects); i++ {
		if w >= projects[i].capital {
			heap.Push(maxHeap, projects[i].profit)
		} else if k > 0 && maxHeap.Len() > 0 {
			k--
			i--
			w += heap.Pop(maxHeap).(int)
		}
	}

	for maxHeap.Len() > 0 && k > 0 {
		w += heap.Pop(maxHeap).(int)
		k--
	}

	return w
}
