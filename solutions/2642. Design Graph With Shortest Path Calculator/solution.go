import "math"

type Graph struct {
	graph map[int]map[int]int
	n     int
}

func Constructor(n int, edges [][]int) Graph {
	graph := Graph{make(map[int]map[int]int), n}
	for _, edge := range edges {
		graph.AddEdge(edge)
	}
	return graph
}

func (this *Graph) AddEdge(edge []int) {
	p := edge[0]
	q := edge[1]
	c := edge[2]
	if _, ok := this.graph[p]; !ok {
		this.graph[p] = make(map[int]int)
	}
	this.graph[p][q] = c
}

func (this *Graph) ShortestPath(node1 int, node2 int) int {
	dist := make([]int, this.n)
	for i := 0; i < this.n; i++ {
		dist[i] = math.MaxInt
	}
	dist[node1] = 0
	heap := heapq_constructor()
	heap.heappush([]int{0, node1})
	for !heap.isEmpty() {
		pair := heap.heappop()
		d := pair[0]
		node := pair[1]
		if d > dist[node] {
			continue
		}
		if node == node2 {
			return d
		}
		if _, ok := this.graph[node]; !ok {
			continue
		}
		for nei, c := range this.graph[node] {
			val := d + c
			if val < dist[nei] {
				dist[nei] = val
				heap.heappush([]int{val, nei})
			}
		}
	}
	return -1
}

/**
 * Your Graph object will be instantiated and called as such:
 * obj := Constructor(n, edges);
 * obj.AddEdge(edge);
 * param_2 := obj.ShortestPath(node1,node2);
 */

type heapq struct {
	heap [][]int
}

func heapq_constructor() heapq {
	return heapq{[][]int{}}
}

func (this *heapq) isEmpty() bool {
	return len(this.heap) == 0
}

func (this *heapq) heappush(value []int) {
	this.heap = append(this.heap, value)
	last_idx := len(this.heap) - 1
	this.shift_up(last_idx)
}

func (this *heapq) heappop() []int {
	last_idx := len(this.heap) - 1
	this.swap(0, last_idx)
	item := this.heap[last_idx]
	this.heap = this.heap[:last_idx]
	this.shift_down(0)
	return item
}

func (this *heapq) swap(i, j int) {
	this.heap[i], this.heap[j] = this.heap[j], this.heap[i]
}

func (this *heapq) shift_up(idx int) int {
	if idx == 0 {
		return -1
	}
	parrent_idx := (idx - 1) / 2
	if this.heap[idx][0] > this.heap[parrent_idx][0] {
		return -1
	} else if this.heap[parrent_idx][0] > this.heap[idx][0] {
		this.swap(idx, parrent_idx)
		return this.shift_up(parrent_idx)
	} else if this.heap[parrent_idx][1] > this.heap[idx][1] {
		this.swap(idx, parrent_idx)
		return this.shift_up(parrent_idx)
	}
	return -1
}

func (this *heapq) shift_down(idx int) int {
	left_idx := (2 * idx) + 1
	right_idx := (2 * idx) + 2
	var child_idx int
	if left_idx >= len(this.heap) {
		return -1
	} else if right_idx >= len(this.heap) {
		child_idx = left_idx
	} else if this.heap[left_idx][0] < this.heap[right_idx][0] {
		child_idx = left_idx
	} else if this.heap[right_idx][0] < this.heap[left_idx][0] {
		child_idx = right_idx
	} else if this.heap[left_idx][1] < this.heap[right_idx][1] {
		child_idx = left_idx
	} else if this.heap[right_idx][1] < this.heap[left_idx][1] {
		child_idx = right_idx
	}
	if this.heap[child_idx][0] > this.heap[idx][0] {
		return -1
	} else if this.heap[idx][0] > this.heap[child_idx][0] {
		this.swap(idx, child_idx)
		return this.shift_down(child_idx)
	} else if this.heap[idx][1] > this.heap[child_idx][1] {
		this.swap(idx, child_idx)
		return this.shift_down(child_idx)
	}
	return -1
}
