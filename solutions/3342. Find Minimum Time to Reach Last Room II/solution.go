import (
	"container/heap"
	"math"
)

type State struct {
	x, y, dis, travelTime int
}

type StateHeap []State

func (h StateHeap) Len() int {
	return len(h)
}

func (h StateHeap) Less(i, j int) bool {
	return h[i].dis < h[j].dis
}

func (h StateHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *StateHeap) Push(x interface{}) {
	*h = append(*h, x.(State))
}

func (h *StateHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func minTimeToReach(moveTime [][]int) int {
	n, m := len(moveTime), len(moveTime[0])
	d := make([][]int, n)
	v := make([][]bool, n)
	for i := range d {
		d[i] = make([]int, m)
		v[i] = make([]bool, m)
		for j := range d[i] {
			d[i][j] = math.MaxInt32
		}
	}

	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	d[0][0] = 0
	h := &StateHeap{}
	heap.Push(h, State{0, 0, 0, 0})

	for h.Len() > 0 {
		s := heap.Pop(h).(State)
		if v[s.x][s.y] {
			continue
		}
		v[s.x][s.y] = true
		travelTime := s.travelTime
		if travelTime == 1 {
			travelTime = 2
		} else {
			travelTime = 1
		}
		for _, dir := range dirs {
			nx, ny := s.x+dir[0], s.y+dir[1]
			if nx < 0 || nx >= n || ny < 0 || ny >= m {
				continue
			}
			dist := max(d[s.x][s.y], moveTime[nx][ny]) + travelTime
			if d[nx][ny] > dist {
				d[nx][ny] = dist
				heap.Push(h, State{nx, ny, dist, travelTime})
			}
		}
	}

	return d[n-1][m-1]
}
