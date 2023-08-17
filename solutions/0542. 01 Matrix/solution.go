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

var directions = [][]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func updateMatrix(mat [][]int) [][]int {
	if mat == nil || len(mat) == 0 || len(mat[0]) == 0 {
		return [][]int{}
	}

	n, m := len(mat), len(mat[0])
	q := NewQueue[[]int]()

	for i := range mat {
		for j := range mat[i] {
			if mat[i][j] == 0 {
				q.Push([]int{i, j})
			} else {
				mat[i][j] = m * n
			}
		}
	}

	for !q.IsEmpty() {
		cell := q.Pop()

		for _, dir := range directions {
			i, j := cell[0]+dir[0], cell[1]+dir[1]
			if i >= 0 && i < n && j >= 0 && j < m && mat[i][j] > mat[cell[0]][cell[1]]+1 {
				q.Push([]int{i, j})
				mat[i][j] = mat[cell[0]][cell[1]] + 1
			}
		}
	}

	return mat
}
