import "container/heap"

type Point struct {
	x, y, cost int
}

type MinHeap []Point

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Point))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minCost(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dist := make([][]int, m)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = 1e9
		}
	}
	dist[0][0] = 0

	directions := []struct{ dx, dy int }{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	h := &MinHeap{}
	heap.Push(h, Point{0, 0, 0}) // Push starting point (0, 0) with 0 cost

	for h.Len() > 0 {
		cell := heap.Pop(h).(Point)
		x, y, cost := cell.x, cell.y, cell.cost

		if cost > dist[x][y] {
			continue
		}

		for i := 0; i < 4; i++ {
			newX, newY := x+directions[i].dx, y+directions[i].dy
			if newX < 0 || newY < 0 || newX >= m || newY >= n {
				continue
			}

			newCost := cost
			if grid[x][y] != i+1 {
				newCost++
			}

			if newCost < dist[newX][newY] {
				dist[newX][newY] = newCost
				heap.Push(h, Point{newX, newY, newCost})
			}
		}
	}

	return dist[m-1][n-1]
}
