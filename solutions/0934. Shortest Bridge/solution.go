type Queue struct {
	x []int
	y []int
}

func NewQueue() *Queue {
	return &Queue{
		x: make([]int, 0),
		y: make([]int, 0),
	}
}

func (q *Queue) Clear() {
	q.x = []int{}
	q.y = []int{}
}

func (q *Queue) Push(x int, y int) {
	q.x = append(q.x, x)
	q.y = append(q.y, y)
}

func (q *Queue) Peek() (int, int) {
	return q.x[0], q.y[0]
}

func (q *Queue) Pop() (int, int) {
	defer func() {
		q.x = q.x[1:]
		q.y = q.y[1:]
	}()

	return q.Peek()
}

func (q *Queue) Size() int {
	return len(q.x)
}

func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}

type coord struct {
	x int
	y int
}

func validCoord(grid [][]int, x, y int) bool {
	return x >= 0 && y >= 0 && x < len(grid) && y < len(grid[0]) && grid[x][y] == 1
}

func shortestBridge(grid [][]int) int {
	q := NewQueue()
	islands := make([][]coord, 2)
	islands[0] = make([]coord, 0)
	islands[1] = make([]coord, 0)

	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[i]))
	}

	bfs := func(x int, y int, n int) {
		q.Clear()
		q.Push(x, y)

		for !q.IsEmpty() {
			cx, cy := q.Pop()

			if !validCoord(grid, cx-1, cy) ||
				!validCoord(grid, cx+1, cy) ||
				!validCoord(grid, cx, cy-1) ||
				!validCoord(grid, cx, cy+1) {
				islands[n] = append(islands[n], coord{
					x: cx,
					y: cy,
				})
			}

			if validCoord(grid, cx-1, cy) && !visited[cx-1][cy] {
				visited[cx-1][cy] = true
				q.Push(cx-1, cy)
			}
			if validCoord(grid, cx+1, cy) && !visited[cx+1][cy] {
				visited[cx+1][cy] = true
				q.Push(cx+1, cy)
			}
			if validCoord(grid, cx, cy-1) && !visited[cx][cy-1] {
				visited[cx][cy-1] = true
				q.Push(cx, cy-1)
			}
			if validCoord(grid, cx, cy+1) && !visited[cx][cy+1] {
				visited[cx][cy+1] = true
				q.Push(cx, cy+1)
			}
		}
	}

	islandNumber := 0
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == 1 && !visited[x][y] {
				bfs(x, y, islandNumber)
				islandNumber++
			}
		}
	}

	min := -1
	for _, c1 := range islands[0] {
		for _, c2 := range islands[1] {
			res := Abs(c1.x-c2.x) + Abs(c1.y-c2.y)
			if min == -1 || res < min {
				min = res
			}
		}
	}

	return min - 1
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
