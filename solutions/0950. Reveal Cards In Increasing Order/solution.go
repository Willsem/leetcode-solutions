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

func deckRevealedIncreasing(deck []int) []int {
	q := NewQueue[int]()
	for i := range deck {
		q.Push(i)
	}

	sort.Ints(deck)

	res := make([]int, len(deck))
	for i := range deck {
		res[q.Pop()] = deck[i]

		if !q.IsEmpty() {
			q.Push(q.Pop())
		}
	}

	return res
}
