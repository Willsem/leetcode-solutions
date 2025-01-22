type Queue[T any] struct {
	data []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		data: make([]T, 0),
	}
}

func (q *Queue[T]) Clear() {
	q.data = []T{}
}

func (q *Queue[T]) Push(el T) {
	q.data = append(q.data, el)
}

func (q *Queue[T]) Peek() T {
	return q.data[0]
}

func (q *Queue[T]) Pop() T {
	defer func() {
		q.data = q.data[1:]
	}()

	return q.Peek()
}

func (q *Queue[T]) Size() int {
	return len(q.data)
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Size() == 0
}

type cell struct {
	x, y int
}

var (
	dx = []int{+0, +0, +1, -1}
	dy = []int{+1, -1, +0, +0}
)

func highestPeak(isWater [][]int) [][]int {
	n, m := len(isWater), len(isWater[0])

	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, m)
		for j := range result[i] {
			result[i][j] = -1
		}
	}

	q := NewQueue[cell]()

	for x := range n {
		for y := range m {
			if isWater[x][y] == 1 {
				q.Push(cell{x, y})
				result[x][y] = 0
			}
		}
	}

	currHeight := 1
	for !q.IsEmpty() {
		size := q.Size()
		for range size {
			curr := q.Pop()

			for d := range 4 {
				nextX := curr.x + dx[d]
				nextY := curr.y + dy[d]

				if nextX < 0 || nextX >= n || nextY < 0 || nextY >= m {
					continue
				}

				if result[nextX][nextY] == -1 {
					result[nextX][nextY] = currHeight
					q.Push(cell{nextX, nextY})
				}
			}
		}

		currHeight++
	}

	return result
}
