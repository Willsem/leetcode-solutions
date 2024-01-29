type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		data: make([]T, 0),
	}
}

func (s *Stack[T]) Push(el T) {
	s.data = append(s.data, el)
}

func (s *Stack[T]) Pop() T {
	defer func() {
		s.data = s.data[:len(s.data)-1]
	}()

	return s.Peek()
}

func (s *Stack[T]) Peek() T {
	return s.data[len(s.data)-1]
}

func (s *Stack[T]) Size() int {
	return len(s.data)
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Size() == 0
}

type MyQueue struct {
	s1    *Stack[int]
	s2    *Stack[int]
	front int
}

func Constructor() MyQueue {
	return MyQueue{
		s1: NewStack[int](),
		s2: NewStack[int](),
	}
}

func (q *MyQueue) Push(x int) {
	if q.s1.IsEmpty() {
		q.front = x
	}
	q.s1.Push(x)
}

func (q *MyQueue) Pop() int {
	if q.s2.IsEmpty() {
		for !q.s1.IsEmpty() {
			q.s2.Push(q.s1.Pop())
		}
	}
	return q.s2.Pop()
}

func (q *MyQueue) Peek() int {
	if !q.s2.IsEmpty() {
		return q.s2.Peek()
	}

	return q.front
}

func (q *MyQueue) Empty() bool {
	return q.s1.IsEmpty() && q.s2.IsEmpty()
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
