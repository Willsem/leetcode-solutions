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

func minLength(s string) int {
	stack := NewStack[rune]()

	for _, c := range s {
		if stack.IsEmpty() {
			stack.Push(c)
			continue
		}

		if c == 'B' && stack.Peek() == 'A' || c == 'D' && stack.Peek() == 'C' {
			stack.Pop()
		} else {
			stack.Push(c)
		}
	}

	return stack.Size()
}
