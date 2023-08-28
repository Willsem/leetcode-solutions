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

type MyStack struct {
	q1 *Queue[int]
	q2 *Queue[int]
}

func Constructor() MyStack {
	return MyStack{
		q1: NewQueue[int](),
		q2: NewQueue[int](),
	}
}

func (s *MyStack) Push(x int) {
	s.q2.Push(x)
	for !s.q1.IsEmpty() {
		s.q2.Push(s.q1.Pop())
	}
	for !s.q2.IsEmpty() {
		s.q1.Push(s.q2.Pop())
	}
}

func (s *MyStack) Pop() int {
	return s.q1.Pop()
}

func (s *MyStack) Top() int {
	return s.q1.Peek()
}

func (s *MyStack) Empty() bool {
	return s.q1.IsEmpty()
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
