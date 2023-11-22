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

type coord struct {
	x, y int
}

func findDiagonalOrder(nums [][]int) []int {
	q := NewQueue[coord]()
	q.Push(coord{0, 0})

	ans := make([]int, 0)

	for !q.IsEmpty() {
		curr := q.Pop()

		ans = append(ans, nums[curr.x][curr.y])

		if curr.y == 0 && curr.x+1 < len(nums) {
			q.Push(coord{curr.x + 1, curr.y})
		}

		if curr.y+1 < len(nums[curr.x]) {
			q.Push(coord{curr.x, curr.y + 1})
		}
	}

	return ans
}
