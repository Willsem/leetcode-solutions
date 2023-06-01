type point struct {
	x     int
	y     int
	depth int
}

type Queue struct {
	data []point
}

func NewQueue() *Queue {
	return &Queue{
		data: make([]point, 0),
	}
}

func (q *Queue) Clear() {
	q.data = []point{}
}

func (q *Queue) Push(el point) {
	q.data = append(q.data, el)
}

func (q *Queue) Peek() point {
	return q.data[0]
}

func (q *Queue) Pop() point {
	defer func() {
		q.data = q.data[1:]
	}()

	return q.Peek()
}

func (q *Queue) Size() int {
	return len(q.data)
}

func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}

var (
	dxarr = []int{-1, 0, 1}
	dyarr = []int{-1, 0, 1}
)

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])

	if grid[0][0] == 1 || grid[n-1][m-1] == 1 {
		return -1
	}

	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, m)
	}
	q := NewQueue()

	q.Push(point{0, 0, 1})
	matrix[0][0] = 1

	for !q.IsEmpty() {
		curr := q.Pop()

		for _, dx := range dxarr {
			for _, dy := range dyarr {
				if dx == 0 && dy == 0 {
					continue
				}

				x := curr.x + dx
				y := curr.y + dy

				if x < 0 || x >= n {
					continue
				}

				if y < 0 || y >= m {
					continue
				}

				newD := curr.depth + 1
				if grid[x][y] == 0 && (matrix[x][y] == 0 || matrix[x][y] > newD) {
					q.Push(point{x, y, newD})
					matrix[x][y] = newD
				}
			}
		}
	}

	res := matrix[n-1][m-1]
	if res == 0 {
		return -1
	}
	return res
}
