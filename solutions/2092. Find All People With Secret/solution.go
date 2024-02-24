import "container/heap"

type Meeting struct {
	person int
	time   int
}

type MinHeap []Meeting

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].time < h[j].time }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Meeting)) }

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	graph := make([][]Meeting, n)
	for _, meeting := range meetings {
		p1, p2, t := meeting[0], meeting[1], meeting[2]

		graph[p1] = append(graph[p1], Meeting{p2, t})
		graph[p2] = append(graph[p2], Meeting{p1, t})
	}

	visited := make([]bool, n)
	q := MinHeap{{0, 0}, {firstPerson, 0}}

	for q.Len() > 0 {
		f := heap.Pop(&q)
		person, meetTime := f.(Meeting).person, f.(Meeting).time

		if visited[person] {
			continue
		}
		visited[person] = true

		for _, v := range graph[person] {
			if v.time >= meetTime && !visited[v.person] {
				heap.Push(&q, v)
			}
		}
	}

	ans := make([]int, 0)
	for i := range visited {
		if visited[i] {
			ans = append(ans, i)
		}
	}
	return ans
}
