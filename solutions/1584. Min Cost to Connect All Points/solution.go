import (
	"container/heap"
	"math"
)

type UnionFind struct {
	parent, rank []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &UnionFind{parent, rank}
}

func (uf *UnionFind) find(u int) int {
	if uf.parent[u] == u {
		return u
	}
	uf.parent[u] = uf.find(uf.parent[u])
	return uf.parent[u]
}

func (uf *UnionFind) union(u, v int) bool {
	u, v = uf.find(u), uf.find(v)
	if u == v {
		return false
	}
	if uf.rank[u] > uf.rank[v] {
		u, v = v, u
	}
	uf.parent[u] = v
	if uf.rank[u] == uf.rank[v] {
		uf.rank[v]++
	}
	return true
}

func manhattanDistance(p1, p2 []int) int {
	return int(math.Abs(float64(p1[0]-p2[0])) + math.Abs(float64(p1[1]-p2[1])))
}

type Edge struct {
	distance, u, v int
}

type MinHeap []Edge

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].distance < h[j].distance }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Edge)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func minCostConnectPoints(points [][]int) int {
	n := len(points)
	uf := NewUnionFind(n)

	edges := &MinHeap{}
	heap.Init(edges)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			distance := manhattanDistance(points[i], points[j])
			heap.Push(edges, Edge{distance, i, j})
		}
	}

	mstWeight := 0
	mstEdges := 0

	for edges.Len() > 0 {
		edge := heap.Pop(edges).(Edge)
		if uf.union(edge.u, edge.v) {
			mstWeight += edge.distance
			mstEdges++
			if mstEdges == n-1 {
				break
			}
		}
	}

	return mstWeight
}
