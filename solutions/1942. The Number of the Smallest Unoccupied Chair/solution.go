import (
	"container/heap"
	"sort"
)

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

type EventHeap [][2]int

func (h EventHeap) Len() int           { return len(h) }
func (h EventHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h EventHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *EventHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *EventHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func smallestChair(times [][]int, targetFriend int) int {
	n := len(times)

	arrivals := make([][2]int, n)
	for i := 0; i < n; i++ {
		arrivals[i] = [2]int{times[i][0], i}
	}

	sort.Slice(arrivals, func(i, j int) bool {
		return arrivals[i][0] < arrivals[j][0]
	})

	availableChairs := &IntHeap{}
	heap.Init(availableChairs)
	for i := 0; i < n; i++ {
		heap.Push(availableChairs, i)
	}

	leavingQueue := &EventHeap{}
	heap.Init(leavingQueue)

	for _, arrival := range arrivals {
		arrivalTime, friendIndex := arrival[0], arrival[1]

		for leavingQueue.Len() > 0 && (*leavingQueue)[0][0] <= arrivalTime {
			heap.Push(availableChairs, heap.Pop(leavingQueue).([2]int)[1])
		}

		chair := heap.Pop(availableChairs).(int)

		if friendIndex == targetFriend {
			return chair
		}

		heap.Push(leavingQueue, [2]int{times[friendIndex][1], chair})
	}

	return -1
}
