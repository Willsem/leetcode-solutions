import (
	"container/heap"
	"math"
	"sort"
)

type maxHeap []int

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h *maxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() any {
	x := (*h)[h.Len()-1]
	*h = (*h)[0 : h.Len()-1]
	return x
}

type worker struct {
	quality int
	wage    int
}

func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
	workers := make([]worker, 0, len(quality))
	for i := range quality {
		workers = append(workers, worker{quality[i], wage[i]})
	}
	sort.Slice(workers, func(i, j int) bool {
		return workers[i].wage*workers[j].quality < workers[j].wage*workers[i].quality
	})

	minCost := math.MaxFloat64
	h := make(maxHeap, 0)
	qSum := 0
	for _, worker := range workers {
		heap.Push(&h, worker.quality)
		qSum += worker.quality
		if h.Len() > k {
			qSum -= heap.Pop(&h).(int)
		}

		cost := float64(worker.wage*qSum) / float64(worker.quality)
		if h.Len() == k && cost < minCost {
			minCost = cost
		}
	}

	return minCost
}
