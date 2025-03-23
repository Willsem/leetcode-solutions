import (
	"container/heap"
	"math"
)

const MOD = 1_000_000_007

type HeapElement struct {
	time, node int
}

type Heap []HeapElement

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].time < h[j].time }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(HeapElement))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func countPaths(n int, roads [][]int) int {
	adj := make([][]HeapElement, n)
	for _, road := range roads {
		u, v, t := road[0], road[1], road[2]
		adj[u] = append(adj[u], HeapElement{t, v})
		adj[v] = append(adj[v], HeapElement{t, u})
	}

	shortestTime := make([]int, n)
	cnt := make([]int, n)
	for i := range shortestTime {
		shortestTime[i] = math.MaxInt64
	}
	shortestTime[0] = 0
	cnt[0] = 1

	pq := &Heap{{0, 0}}
	heap.Init(pq)

	for pq.Len() > 0 {
		top := heap.Pop(pq).(HeapElement)
		time, node := top.time, top.node

		if time > shortestTime[node] {
			continue
		}

		for _, neighbor := range adj[node] {
			nbr, rtime := neighbor.node, neighbor.time
			newTime := time + rtime

			if newTime < shortestTime[nbr] {
				shortestTime[nbr] = newTime
				cnt[nbr] = cnt[node]
				heap.Push(pq, HeapElement{newTime, nbr})
			} else if newTime == shortestTime[nbr] {
				cnt[nbr] = (cnt[nbr] + cnt[node]) % MOD
			}
		}
	}

	return cnt[n-1]
}
